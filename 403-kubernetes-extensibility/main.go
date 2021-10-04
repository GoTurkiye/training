package gophercon_turkiye_2021_hands_on

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const RequiredLabel string = "gophercon.turkiye/available"

var universalDeserializer = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()

func Validate(rw http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		} else {
			fmt.Println(fmt.Errorf("error reading body"))
			http.Error(rw, "Error reading body", http.StatusBadRequest)
			return
		}
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		fmt.Println(fmt.Errorf("contentType=%s, expect application/json", contentType))
		http.Error(rw, "Error parsing body", http.StatusBadRequest)
		return
	}

	fmt.Println(fmt.Sprintf("handling request: %s", body))

	requestedAdmissionReview := v1.AdmissionReview{}
	if _, _, err := universalDeserializer.Decode(body, nil, &requestedAdmissionReview); err != nil {
		fmt.Println(fmt.Sprintf("Error parsing body %v", err))
		http.Error(rw, "Error parsing body", http.StatusBadRequest)
		return
	}

	responseAdmissionReview := v1.AdmissionReview{}
	responseAdmissionReview.TypeMeta = requestedAdmissionReview.TypeMeta
	responseAdmissionReview.Response = validateLabel(requestedAdmissionReview.Request)
	responseAdmissionReview.Response.UID = requestedAdmissionReview.Request.UID

	fmt.Println(fmt.Sprintf("sending response: %v", responseAdmissionReview))
	respBytes, err := json.Marshal(responseAdmissionReview)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		http.Error(rw, fmt.Sprintf("could not marshal response, err: %v", err), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	if _, err := rw.Write(respBytes); err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		http.Error(rw, fmt.Sprintf("could not write response, err: %v", err), http.StatusInternalServerError)
		return
	}
}

func validateLabel(r *v1.AdmissionRequest) *v1.AdmissionResponse {
	response := v1.AdmissionResponse{}
	var pod corev1.Pod

	if _, _, err := universalDeserializer.Decode(r.Object.Raw, nil, &pod); err != nil {
		response.Allowed = false
		response.Result = &metav1.Status{
			Message: err.Error(),
			Reason:  "could not marshall request object to Pod",
		}
		return &response
	}

	allowed := false
	for k, v := range pod.Labels {
		if k == RequiredLabel && v == "true" {
			allowed = true
			break
		}
	}

	if !allowed {
		response.Result = &metav1.Status{
			Message: fmt.Sprintf("required label \"%s\" should be exist within the Pod labels with the values \"true\" or \"false\"", RequiredLabel),
		}
	}

	response.Allowed = allowed
	return &response
}
