# Prerequisities

Selamlar herkese, bugünkü eğitimimizin mümkün olduğunca interaktif geçmesi adına sizlerden bir kaç aracın kurulumu hazır şekilde sunuma katılmanızı isteyeceğim.

1. **Docker** makinelerimizde kurulu olmalı çünkü bir Go uygulamasını Dockerfile hazırlayarak (farklı yöntemlere de değinebiliriz, Dockerfile zorunlu değil) uygulamamızı bir container image haline getireceğiz. Aşağıdaki linki takip ederek kurulumu gerçekleştirebilirsiniz:

    * https://docs.docker.com/get-docker/

    Eğer her şey yolunda gittiyse `$ docker version` komutunu çalıştırın ve her şeyin başarılı bir şekilde çalıştığından emin olun.

2. Bugün projemizin release süreçleri sırasında [GoReleaser](https://goreleaser.com) adlı araçtan faydalanacağız. Aşağıdaki linki takip ederek kurulumu gerçekleştirebilirsiniz:

    * https://goreleaser.com/install/

    Eğer her şey yolunda gittiyse `$ goreleaser --version` komutunu çalıştırın ve her şeyin başarılı bir şekilde çalıştığından emin olun.

3. Tüm source kodlarımızı Github üzerinden hallediyor olacağız, dolayısıyla eğer Github hesabınız yoksa aşağıdaki linki takip ederek bir GitHub hesabı oluşturabilirsiniz:

    * https://github.com/signup

