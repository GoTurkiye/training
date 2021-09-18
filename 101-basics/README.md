# [101] Basics

Bu materyal Go dilinin temel pratiklerini, yazım stilini anlatan örnekler içerir.

## Syntax

Go dilin diğer dillerden farklı olarak `;` kullanımı belirli ayraçlar hariç zorunlu kılınmamıştır. 
Yine çoğu dilden farklı olarak `null` yerine `nil` kullanılır. 
Bunların yanında en belirgin diğer fark ise `exception`dır. 
Go dilinde `throwable` bir `exception` yapısı yoktur.
`exception`ların yerini Go'da `error` interfacei alır ve hata durumları bu interface üzerinden yönetilir.

Dilin temelinde 25 keyword bulunur. Bunlar:
* break     
* case      
* chan      
* const     
* continue  
* default      
* defer        
* else         
* fallthrough  
* for          
* func    
* go      
* goto    
* if      
* import  
* interface  
* map        
* package    
* range      
* return   
* select
* struct
* switch
* type
* var