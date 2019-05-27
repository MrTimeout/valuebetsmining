$(document).ready(function () {
  $("#menu").click(function () {
    $("nav ul li").toggleClass("oculto");
  });
  $(window).resize(function () {
    $(".oculto").removeClass();
  });
  $("#contacto").click(function () {
    $(".login").toggleClass("visto");
  });
  $("#x").click(function () {
    $(".login").toggleClass("visto");
  });
  $(window).resize(function () {
    $(".login").removeClass("visto");
  });
  $("h1").click(function () {
    var div = "div." + this.className;
    $(div).toggle();
});

/*----------------------Seleccionar--------------------*/

  $(function () {
    seleccionPaises();
    seleccionDivisiones();
    seleccionLocales();
    seleccionVisitantes();
  });

  $("select[name='Paises']").change(function () {
    seleccionDivisiones();
  });

  $("select[name='Division']").change(function () {
    seleccionLocales();
  });

  $("select[name='Local']").change(function () {
    seleccionVisitantes();
  });

/*----------------Calcular--------------------*/
  $('#calcular').click(function () {
    var num1 = $('#localA').val();
    var num2 = $('#localB').val();
    var num3 = 'resultadolocal';
    var num4 = 'kellylocal';
    calcula(num1, num2, num3, num4);
    var a = $('#empateA').val();
    var b = $('#empateB').val();
    var c = 'resultadoempate';
    var d = 'kellyempate';
    calcula(a, b, c, d);
    var x = $('#visitanteA').val();
    var y = $('#visitanteB').val();
    var z = 'resultadovisitante';
    var s = 'kellyvisitante';
    calcula(x, y, z, s);
  });

/*-------------FormContacto---------------------*/
  

  $('#btnSend').click(function(){

    var errores = '';

    // Validado Nombre ==============================
    if( $('#names').val() == '' ){
        errores += '<p>Escriba un nombre</p>';
        $('#names').css("border-bottom-color", "#F14B4B")
    } else{
        $('#names').css("border-bottom-color", "#d1d1d1")
    }

    // Validado Correo ==============================
    if( $('#email').val() == '' ){
        errores += '<p>Ingrese un correo</p>';
        $('#email').css("border-bottom-color", "#F14B4B")
    } else{
        $('#email').css("border-bottom-color", "#d1d1d1")
    }

    // Validado Mensaje ==============================
    if( $('#mensaje').val() == '' ){
        errores += '<p>Escriba un mensaje</p>';
        $('#mensaje').css("border-bottom-color", "#F14B4B")
    } else{
        $('#mensaje').css("border-bottom-color", "#d1d1d1")
    }
/*
    // ENVIANDO MENSAJE ============================
    if( errores == '' == false){
        var mensajeModal = '<div class="modal_wrap">'+
                                '<div class="mensaje_modal">'+
                                    '<h3>Errores encontrados</h3>'+
                                    errores+
                                    '<span id="btnClose">Cerrar</span>'+
                                '</div>'+
                            '</div>'

        $('body').append(mensajeModal);
    }
*/
    // CERRANDO MODAL ==============================
    $('#btnClose').click(function(){
        $('.modal_wrap').remove();
    });
});


});


/*---------------------------------------------------SELECCION---------------------------*/




var paises = ["Alemania", "Francia", "España"];
var divisiones = ["Primera", "Segunda"];
var locales = ["Barcelona", "Madrid", "Celta"];
var visitantes = ["Girona", "Valencia", "Sevilla"];

function seleccionPaises() {
  $("select[name='Paises']").empty();
  for (var i = 0; i < paises.length; i++) {
    var option = $("<option></option>");
    $(option).html(paises[i]);
    $("select[name='Paises']").append(option);
  }
}

function seleccionDivisiones() {
  $("select[name='Division']").empty();
  for (var i = 0; i < divisiones.length; i++) {
    var option = $("<option></option>");
    $(option).html(divisiones[i]);
    $("select[name='Division']").append(option);
  }
  $("select[name='Local']").empty();
  $("select[name='Visitante']").empty();
}

function seleccionLocales() {
  $("select[name='Local']").empty();
  for (var i = 0; i < locales.length; i++) {
    var option = $("<option></option>");
    $(option).html(locales[i]);
    $("select[name='Local']").append(option);
  }
}

function seleccionVisitantes() {
  $("select[name='Visitante']").empty();
  for (var i = 0; i < visitantes.length; i++) {
    var option = $("<option></option>");
    $(option).html(visitantes[i]);
    $("select[name='Visitante']").append(option);
  }
}


/*----------------------GRAFICAS-----------------------------*/
google.charts.load('current', {
  'packages': ['bar']
});

google.charts.load("current", {
  packages: ["corechart"]
});

var options = {
  chart: {
    color: 'grey'
  },
  legend: {
    position: 'none',
    textStyle: {
      color: 'black'
    }
  },
  hAxis: {
    textStyle: {
      color: 'gray'
    },
    gridlines: {
      color: 'gray'
    }
  },
  vAxis: {
    textStyle: {
      color: 'gray'
    },
    gridlines: {
      color: 'gray'
    }
  },
  colors: ['green', '#C55B34'],
  backgroundColor: '#192126',

};

var options2 = {
  is3D: true,
  colors: ['green', 'blue', '#C55B34'],
  backgroundColor: '#192126',
  textStyle: {
    color: 'black'
  },
  legend: 'none',
};

google.charts.setOnLoadCallback(drawChart);
google.charts.setOnLoadCallback(tablaGoles);
google.charts.setOnLoadCallback(tablaRachaG);
google.charts.setOnLoadCallback(tablaPorcentaje);


function drawChart() {
  var resultados = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['GANADOS', 5, 2],
    ['EMPATADOS', 2, 2],
    ['PERDIDOS', 3, 6]
  ]);
  var chart = new google.charts.Bar(document.getElementById('ultimos'));

  chart.draw(resultados, google.charts.Bar.convertOptions(options));
}

function tablaGoles() {
  var goles = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['A FAVOR', 5, 2],
    ['EN CONTRA', 8, 2]
  ]);
  var chart = new google.charts.Bar(document.getElementById('goles'));

  chart.draw(goles, google.charts.Bar.convertOptions(options));
}

function tablaRachaG() {
  var ganados = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['GANADOS', 5, 4],
    ['INVICTO', 8, 1]
  ]);
  var chart = new google.charts.Bar(document.getElementById('racha'));

  chart.draw(ganados, google.charts.Bar.convertOptions(options));
}

function tablaPorcentaje() {
  var posibilidades = google.visualization.arrayToDataTable([
    ['%', 'Probabilidades'],
    ['LOCAL', 7],
    ['EMPATE', 2],
    ['VISITANTE', 1]
  ]);
  var chart = new google.visualization.PieChart(document.getElementById('porcentaje'));
  chart.draw(posibilidades, options2);
}
/**--------------------------------Grafica Accesible--------------------------- **/




/*-------------------------------Calculos----------------------------*/
function calcula(n1, n2, n3, n4) {
  //Almaceno los valores de los inputs
  var primerValor = n1;
  var segundoValor = n2;

  //Condiciona para que acepte solo números usando las expresiones regulares
  if ($.isNumeric(primerValor) && $.isNumeric(segundoValor) && parseFloat(primerValor) < 1 && parseFloat(segundoValor) > 1) {
    var resultado = parseFloat(primerValor) * parseFloat(segundoValor);
    //Muestro el resultado
    $('#' + n3).val(resultado.toFixed(3));
    if (resultado > 1.1) {
      $('#' + n3).css({
        "background-color": "green"
      });
      $('#' + n4).val((n1 - (1 - n1) / (n2 - 1)).toFixed(3));
      $('#' + n4).css({
        "background-color": "green"
      });
    } else {
      $('#' + n3).css({
        "background-color": "gray"
      });
      $('#' + n4).val(0);
      $('#' + n4).css({
        "background-color": "gray"
      });
    }
  } else {
    $('#' + n3).val('xxx');
    $('#' + n3).css({
      "background-color": "red"
    });
    $('#' + n4).val('xxx');
    $('#' + n4).css({
      "background-color": "red"
    });
  }
};