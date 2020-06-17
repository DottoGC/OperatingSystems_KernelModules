
$(document).ready(function() {

    let procInfo = null;
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    const init = {
        method: 'GET',
        headers
    };


    const init2 = {
        method: 'GET',
    };

    $('#overlayProcs').fadeOut(3000,function(){
        $('#divProcs').fadeIn(2000);
    });

    getProcsInfo();
    var table;

    setTimeout(function(){


        var cardProcs = document.getElementById("cardProcesos");
        cardProcs.innerHTML = "<br> Procesos en ejecución: "+ procInfo.Ejecucion +"</br>" +
            " <br>Procesos suspendidos: "+ procInfo.Suspendidos +"</br>" +
            " <br>Procesos detenidos: "+ procInfo.Detenidos +"</br>" +
            " <br>Procesos Zombie: "+ procInfo.Zombie +"</br>" +
            " <br>Total de procesos: "+ procInfo.Total +"</br>" ;

        table = $('#dataTableProcs');

      table = table.DataTable( {
        data: procInfo.Procesos,

        columns: [
            { data:  "Pid"  },
            { data:  "Nombre"  },
            { data:  "Usuario"  },
            { data: "Estado"  },
            { data:  "Porcentaje"  },
            {
                data: "Matar", // can be null or undefine
            }
    ],

    } );


        $('#dataTableProcs tbody').on( 'click', 'button', function () {
            var data = table.row($(this).parents('tr')).data();
            fetch('http://localhost:8080/kill?keys='+data.Pid, init2)
                .catch((e) => {
                    console.log("ERROR: " + e.toString());
                });

              setTimeout(function(){
                  location.reload();
              }, 3000);
            alert( "RIP proceso: " +data.Nombre);
        } );

    }, 8000);





    function getProcsInfo(){

        fetch('http://localhost:8080/procs', init)
            .then(response => response.json())
            .then(data => {
                procInfo = data
                // text is the response body
            })
            .catch((e) => {
                console.log("ERROR: " + e.toString());
            });
    }

});