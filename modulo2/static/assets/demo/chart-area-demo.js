Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';


$(document).ready(function() {

  let memInfo = null;
  let porcentaje = [];
  const headers = new Headers();
  headers.append('Content-Type', 'application/json');

  const init = {
    method: 'GET',
    headers
  };

  fetch('http://localhost:8080/memoria', init)
      .then(response => response.json())
      .then(data => {
        memInfo = data
        // text is the response body
      })
      .catch((e) => {
        console.log("ERROR: " + e.toString());
      });

  debugger
  porcentaje.push(memInfo.porcentaje)
  console.log(porcentaje[0])
});

