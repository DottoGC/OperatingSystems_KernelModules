
$(document).ready(function() {

    let procInfo = null;
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    const init = {
        method: 'GET',
        headers
    };

    $('#overlayProcsArbol').fadeOut(3000,function(){
        $('#divProcsArbol').fadeIn(2000);
    });


    fetch('http://localhost:8080/procsArbol', init)
            .then(response => response.json())
            .then(data => {
                procInfo = data
                // text is the response body
            })
            .catch((e) => {
                console.log("ERROR: " + e.toString());
            });


    setTimeout(function(){

        var table = new Tabulator("#dataGridProcs", {
            height:"311px",
            data:procInfo,
            dataTree:true,
            dataTreeStartExpanded:true,
            columns:[
                {title:"Pid", field:"Pid", width:200, responsive:0}, //never hide this column
                {title:"Nombre", field:"Nombre", width:200},
                {title:"Usuario", field:"Usuario", width:200},
                {title:"Estado", field:"Estado", width:200},

             //   {title:"Hijos", field:"Hijos", width:150, responsive:2}, //hide this column first
             //   {title:"Favourite Color", field:"col", width:150},
             //   {title:"Date Of Birth", field:"dob", hozAlign:"center", sorter:"date", width:150},
            ],
        });

 }, 3000);





});