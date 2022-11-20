const form = document.querySelector('form');

var dataform  = {};
var checkboxes = [] ; 
var filedata ; 





function registerCheckboxes(data){

  data.forEach(function (element) {

 
    if(dataform[element] == 'on'){
      checkboxes.push(element)
    }

  })

}




function localization(date){

  new Intl.DateTimeFormat('id', { dateStyle: 'full', timeStyle: 'long' }).format(new Date(date))

}

function distance(startmilis , endmilis){

  var date1 = new Date(startmilis);
var date2 = new Date(endmilis);
var diffDays = date2.getDate() - date1.getDate();

if(diffDays >= 28){
  return `${(diffDays - (diffDays % 28))/28} Bulan ${diffDays % 28} hari`
}else{
  return `${diffDays} Hari`
}


  

}

function render(object){

  efile = filedata


document.querySelector('.projectlist').innerHTML += `

<div class="project-item">
<div class="project-intro">
    <div class="project-image"><img src="${efile}" alt="" srcset=""></div>
    <div class="project-description">
        <h1>${object.projectName}</h1>
        <p class="duration">${distance(object.startDate , object.endDate)}</p>
    <p>${object.content}</p>
    <div class="project-stack">
        <img class="svg" src="../assets/logo-nodejs.svg" alt="" srcset="">
        <img class="svg" src="../assets/nextjs-icon-svgrepo-com.svg" alt="" srcset="">
        <img class="svg" src="../assets/logo-react.svg" alt="" srcset="">
        
    </div>
    </div>
</div>
<div class="project-controls">
    <div class="btn"><a href="">Details</a></div>
</div>
</div>



`






}


form.addEventListener('submit', (e) => {

    e.preventDefault();

    const formData = new FormData(form);
    const data = Object.fromEntries(formData);
  
  
    // const filteredData = Object.fromEntries(
    //   Object.entries(data).filter(([_, v]) => v)
    // );

    
      // var filedata ; 
      // var filedata  = []; 
      const reader = new FileReader();
      const fileprom = new Promise(function (resolve) {
        reader.readAsDataURL(data.file);
        reader.onloadend = function () {

          resolve(reader.result)


        };


      })

      fileprom.then(data => {
        filedata = data
      })
    

      dataform = {...data ,filedata}; 
       console.log(dataform)
      // form.reset()

      

      //lakukan pendaftaran checkbox


      registerCheckboxes(["nextjs" , "reactjs" , "nodejs" , "typescript"]);

        console.log(dataform.filedata)
      render(dataform)



    


    
});


