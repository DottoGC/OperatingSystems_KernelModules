Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';


$(document).ready(function() {

    let memInfo = null;
    let consumo = [];
    let contador = 0;
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    var chartHtml = document.getElementById("myChartRam").getContext("2d");


    var chartConfig = {
        labels: [],
        datasets: [
            {
                label: "Sessions",
                lineTension: 0.3,
                backgroundColor: "rgba(2,117,216,0.2)",
                borderColor: "rgba(2,117,216,1)",
                pointRadius: 5,
                pointBackgroundColor: "rgba(2,117,216,1)",
                pointBorderColor: "rgba(255,255,255,0.8)",
                pointHoverRadius: 5,
                pointHoverBackgroundColor: "rgba(2,117,216,1)",
                pointHitRadius: 50,
                pointBorderWidth: 2,
                data: [],
            }
        ]
    };

    var options = {
        animation: false,
        //Boolean - If we want to override with a hard coded scale
        scaleOverride: true,
        //** Required if scaleOverride is true **
        //Number - The number of steps in a hard coded scale
        scaleSteps: 10,
        //Number - The value jump in the hard coded scale
        scaleStepWidth: 10,
        //Number - The scale starting value
        scaleStartValue: 0
    };


    var myLineChart = new Chart(chartHtml, {
        type: "line",
        data: chartConfig,
        options: options
    });

    const init = {
        method: 'GET',
        headers
    };

    setInterval(function(){
        getRamInfo();
    }, 5000);



    function getRamInfo(){

        fetch('http://localhost:8080/memoria', init)
            .then(response => response.json())
            .then(data => {
                memInfo = data
                // text is the response body
            })
            .catch((e) => {
                console.log("ERROR: " + e.toString());
            });

        contador++;
        addData(contador, memInfo.consumo)
    }

    function addData(label, data) {
        myLineChart.data.labels.push(label);
        myLineChart.data.datasets.forEach((dataset) => {
            dataset.data.push(data);
        });
        myLineChart.update();
    }
});

