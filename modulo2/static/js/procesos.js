
$(document).ready(function() {

    let procInfo = null;
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    const init = {
        method: 'GET',
        headers
    };

        fetch('http://localhost:8080/procs', init)
            .then(response => response.json())
            .then(data => {
                procInfo = data
                // text is the response body
            })
            .catch((e) => {
                console.log("ERROR: " + e.toString());
            });


    setTimeout(function(){
    var table = $('#dataTableProcs');

    table.DataTable( {
        data: procInfo,
        columns: [
            { data:  "Pid"  },
            { data:  "Nombre"  },
            { data: "Estado"  },
            { data:  "Porcentaje"  }
        ],

    } ); }, 3000);





});