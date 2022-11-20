


/*

references : 

1. https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/formdata_event
2. https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Iterators_and_Generators
3. Itterators to Object

*/

//dataform


var dataform ; 
var datastatus ; 


// kueri dataform
var q = document.querySelector("form")

//eventlistener dari kueri form
q.addEventListener("submit" , (e) => {
    //prevent default
    e.preventDefault();
    //ambil data
    var data = new FormData(q);
    //ambil entri
    data = data.entries();   
    // trigger datatype iterator           
    var obj = data.next();  
    // buat objek penampung
    var dataObj = {} ;   
    // memindahkan iterator kedalam object     
    while(undefined !== obj.value) {
        // pengecekan sederhana (dapat menggunakan fungsi pengecekan)   
        if(obj.value[1] !== ""){
        dataObj[obj.value[0]] = obj.value[1];
         obj = data.next();
         datastatus = true
        }else {
            datastatus = false
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: `${obj.value[0]} tidak boleh kosong`
            })
            break ;
        }
    }


    if(datastatus){
    dataform = dataObj ; 
    Swal.fire(
        'Apakah sudah benar?',
        `Nama : ${dataform.name} Email :  ${dataform.email} Telepon :  ${dataform.tel} Subject :  ${dataform.subject} Pesan :  ${dataform.msg}`,
        'question'
      ).then((result) => {
        if (result.isConfirmed) {
          Swal.fire('Saved!', '', 'success')
        }

        });
    }
    
});


