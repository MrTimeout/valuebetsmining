/*EVENTOS CARGA DOCUMENTO*/
$(document).ready(function () {
  path = "http://localhost:3002/api/v1/"
  recibirPaises()
  /*desplegar menu responsive*/
  $("#menu").click(function () {
    $("nav ul li").toggleClass("oculto");
  });

  /*ocultar menu cambio de resolucion*/
  $(window).resize(function () {
    $(".oculto").removeClass();
  });

  /* carga de formulario inicio sesion  o carga de confirmacion cierre sesion  si esta logeado*/
  $("#contacto").click(function () {
    if ($("#contacto li a img").attr("src") == "recursos/usuario.svg") {
      $(".login").toggleClass("visto");
    } else {
      $(".confirmar").toggleClass("visto");
    }

  });

  /*seleccion de graficas o listas accesibles*/
  $("#graficas").click(function () {
    $("#capa1").toggle();
    $("#capa2").toggle();
  });

  /*cerrar confirmacion cierre sesion*/
  $("#cancelar").click(function () {
    $(".confirmar").toggleClass("visto");
  });

  /*cerrar formulario inicio sesion*/
  $("#x").click(function () {
    $(".login").toggleClass("visto");
  });

  /*cerrar registro si cambia resolucion de pantalla*/
  $(window).resize(function () {
    $(".login").removeClass("visto");
  });

  /*abrir cerrar divs guia */
  $("h1").click(function () {
    var div = "div." + this.className;
    $(div).toggle();
  });


  /*--------Calcular  /Valor   /Kelly--------------------*/

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


  $('#btnSend').click(function () {

    var errores = '';

    // Validado Nombre ==============================
    if ($('#names').val() == '') {
      errores += '<p>Escriba un nombre</p>';
      $('#names').css("border-bottom-color", "#F14B4B")
    } else {
      $('#names').css("border-bottom-color", "#d1d1d1")
    }

    // Validado Correo ==============================
    if ($('#email').val() == '') {
      errores += '<p>Ingrese un correo</p>';
      $('#email').css("border-bottom-color", "#F14B4B")
    } else {
      $('#email').css("border-bottom-color", "#d1d1d1")
    }

    // Validado Mensaje ==============================
    if ($('#mensaje').val() == '') {
      errores += '<p>Escriba un mensaje</p>';
      $('#mensaje').css("border-bottom-color", "#F14B4B")
    } else {
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
    $('#btnClose').click(function () {
      $('.modal_wrap').remove();
    });
  });


  /*------Seleccionar  /  ligas  /  equipos --------------------*/

  $("select[name='Paises']").change(function () {
    $("select#Paises option:contains('Selecciona')").remove()
    recibirDivisiones();
  });

  $("select[name='Division']").change(function () {
    $("select#Division option:contains('Selecciona')").remove()
    recibirLocales();

  });

  $("select[name='Local']").change(function () {
    $("select#Local option:contains('Selecciona')").remove()
    recibirVisitantes();
  });

  $("select[name='Visitante']").change(function () {
    $("select#Visitante option:contains('Selecciona')").remove()
    solicitarAtributos()
  });

});






/*---------------FUNCIONES  DE SELECCION---------------------------*/

/*inicio variables globales*/
var paises = [];
var divisiones = [];
var locales = [];
var visitantes = [];
var atributos;
var pro1;
var proX;
var pro2;
var attHome;
var attAway;

/*solicito paises al servidor*/
function recibirPaises() {
  $.ajax({
    url: path + "countries",
    headers: {
      'Content-Type': 'application/json'
    },
    success: function (result) {
      insertarPaises(result);
    }
  });
};

/*solicito divisiones al servidor*/
function recibirDivisiones() {
  let pais = "" + $('select[name="Paises"] option:selected').text();
  $.ajax({
    url: path + pais + "/divisions",
    success: function (result) {
      insertarDivisiones(result);
    }
  });
};

/*solicito equipos al servidor*/
function recibirLocales() {
  let pais = $('select[name="Paises"] option:selected').text();
  let division = $('select[name="Division"] option:selected').text();
  
  $.ajax({
    url: path + pais + "/" + division + "/teams",
    success: function (result) {
      insertarLocales(result);
    }
  });
};

/*solicito equipos al servidor*/
function recibirVisitantes() {
  let pais = $('select[name="Paises"] option:selected').text();
  let division = $('select[name="Division"] option:selected').text();
  $.ajax({
    url: path + pais + "/" + division + "/teams",
    success: function (result) {
      insertarVisitantes(result);
    }
  });
};

/*inserto paises en el select paises*/
function insertarPaises(p) {
  $("select[name='Paises']").empty();
  let option = $("<option></option>");
  $(option).html("Selecciona");
  $("select[name='Paises']").append(option);
  for (let i = 0; i < p.length; i++) {
    let option = $("<option></option>");
    $(option).html(p[i]);
    $("select[name='Paises']").append(option);
  }
};

/*inserto divisiones en el select divisiones*/
function insertarDivisiones(p) {
  $("select[name='Division']").empty();
  let option = $("<option></option>");
  $(option).html("Selecciona");
  $("select[name='Division']").append(option);
  for (let i = 0; i < p.length; i++) {
    let option = $("<option></option>");
    $(option).html(p[i]);
    $("select[name='Division']").append(option);
  }
  $("select[name='Local']").empty();
  $("select[name='Visitante']").empty();
};

/*inserto equipos en el select locales*/
function insertarLocales(p) {
  $("select[name='Local']").empty();
  let option = $("<option></option>");
  $(option).html("Selecciona");
  $("select[name='Local']").append(option);
  for (let i = 0; i < p.length; i++) {
    let option = $("<option></option>");
    $(option).html(p[i]);
    $("select[name='Local']").append(option);
  }
  $("select[name='Visitante']").empty();
};

/*inserto equipos en el sele      $.ajavisitantes*/
function insertarVisitantes(p) {
  $("select[name='Visitante']").empty();
  let option = $("<option></option>");
  $(option).html("Selecciona");
  $("select[name='Visitante']").append(option);
  for (let i = 0; i < p.length; i++) {
    let option = $("<option></option>");
    $(option).html(p[i]);
    $("select[name='Visitante']").append(option);
  }

};








/*----------------------- Atributos en pagina--------------------------*/


//path+pais + "/" + division + "/" + local + "/properties",
function solicitarAtributos() {
  let pais = "" + $('select[name="Paises"] option:selected').text();
  let division = "" + $('select[name="Division"] option:selected').text();
  let local = "" + $('select[name="Local"] option:selected').text();
  let visitante = "" + $('select[name="Visitante"] option:selected').text();
  attHome=[];
  attAway=[];
  $.ajax({
    url: path + pais + "/" + division + "/"+ local +"/properties?stadium=local",
    success: function (home) {
      $.ajax({
        url: path + pais + "/" + division + "/"+ visitante +"/properties?stadium=away",
        success: function (away) {
          attHome = home;
          attAway= away;
          insertarAtributos()
        }
      });
    }
  });
};

function insertarAtributos() {
  $("#gLocal").text(attHome.Last10WinningMatchs);
  $("#eLocal").text(attHome.Last10TiedingMatchs);
  $("#pLocal").text(attHome.Last10LosingMatchs);
  $("#gVisitante").text(attAway.Last10WinningMatchs);
  $("#eVisitante").text(attAway.Last10TiedingMatchs);
  $("#pVisitante").text(attAway.Last10LosingMatchs);
  $("#gmLocal").text(attHome.Last10GoalsTuckedAmount);
  $("#gmVisitante").text(attAway.Last10GoalsTuckedAmount);
  $("#geLocal").text(attHome.Last10GoalsReceivedAmount);
  $("#geVisitante").text(attAway.Last10GoalsReceivedAmount);
  $("#rLocal").text(attHome.Last10StreackWinning);
  $("#rVisitante").text(attAway.Last10StreackWinning);
  $("#iLocal").text(attHome.Last10StreackNoLosing);
  $("#iVisitante").text(attAway.Last10StreackNoLosing);
  pro1 =0.027 * attHome.Last10WinningMatchs +
    -0.0134 * attAway.Last10WinningMatchs +
    -0.0303 * attAway.Last10TiedingMatchs +
     0.114 * attHome.Last10GoalsTuckedAmount +
    -0.0483 * attHome.Last10GoalsReceivedAmount +
    -0.0882 * attAway.Last10GoalsReceivedAmount +
    -0.0604 * attHome.Last10AverageTuckedGoals +
    0.0888 * attHome.Last10AverageReceivedGoals +
    0.4706;
  pro2 = -0.0132 * attHome.Last10TiedingMatchs +
    0.0358 * attAway.Last10WinningMatchs +
    -0.0215 * attHome.Last10GoalsTuckedAmount +
    0.2181 * attHome.Last10GoalsReceivedAmount +
    -0.0433 * attAway.Last10GoalsReceivedAmount +
    0.0226 * attHome.Last10AverageTuckedGoals +
    -0.1045 * attHome.Last10AverageReceivedGoals +
    0.336;
  proX = -0.0025 * attHome.Last10WinningMatchs +
  0.0049 * attHome.Last10TiedingMatchs +
 -0.0015 * attHome.Last10LosingMatchs +
  0.0049 * attAway.Last10TiedingMatchs +
 -0.0015 * attAway.Last10LosingMatchs +
 -0.6136 * attHome.Last10GoalsTuckedAmount +
  0.5644 * attAway.Last10GoalsTuckedAmount +
 -0.0591 * attAway.Last10GoalsReceivedAmount +
  0.3804;
  $("#proLocal").text((pro1+pro2+proX)/pro1);
  $("#proEmpate").text((pro1+pro2+proX)/proX);
  $("#proVisitante").text((pro1+pro2+proX)/pro2);
  $("#localA").val((pro1+pro2+proX)/pro1);
  $("#emmpateA").val((pro1+pro2+proX)/proX);
  $("#visitanteA").val((pro1+pro2+proX)/pro2);
  insertarGraficas();
}




/*----------------------GRAFICAS-----------------------------*/
google.charts.load('current', {
  'packages': ['bar']
});

google.charts.load("current", {
  packages: ["corechart"]
});


/*----------------------OPCIONES DE GRAFICAS----------------------------*/
var options = {
  chart: {
    color: 'grey'
  },
  legend: {
    position: 'none',
    textStyle: {
      color: 'white'
    }
  },
  hAxis: {
    textStyle: {
      color: 'white'
    },
    gridlines: {
      color: 'gray'
    }
  },
  vAxis: {
    textStyle: {
      color: 'white'
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
    color: 'white'
  },
  legend: 'none',
};


/*---------------CREAR GRAFICAS----------------------*/
function insertarGraficas() {
  google.charts.setOnLoadCallback(drawChart);
  google.charts.setOnLoadCallback(tablaGoles);
  google.charts.setOnLoadCallback(tablaRachaG);
  google.charts.setOnLoadCallback(tablaPorcentaje);
};

/****************** OBTENER GRÁFICAS******************/

/*--------------Grafica RESULTADOS-------------------------*/
function drawChart() {
  var resultados = google.visualization.arrayToDataTable([
    ["Nº", "Local", "Visitante"],
    ["GANADOS", attHome.Last10WinningMatchs, attAway.Last10WinningMatchs],
    ["EMPATADOS", attHome.Last10TiedingMatchs, attAway.Last10TiedingMatchs],
    ["PERDIDOS", attHome.Last10LosingMatchs,attAway.Last10LosingMatchs]
  ]);
  var chart = new google.charts.Bar(document.getElementById('ultimos'));

  chart.draw(resultados, google.charts.Bar.convertOptions(options));
}


/*------------------Grafica GOLES-----------------------------*/
function tablaGoles() {
  var goles = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['A FAVOR', attHome.Last10GoalsTuckedAmount, attAway.Last10GoalsTuckedAmount],
    ['EN CONTRA', attHome.Last10GoalsReceivedAmount, attAway.Last10GoalsReceivedAmount]
  ]);
  var chart = new google.charts.Bar(document.getElementById('goles'));

  chart.draw(goles, google.charts.Bar.convertOptions(options));
}

/*------------------Grafica RACHA-------------------------------------*/
function tablaRachaG() {
  var ganados = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['GANADOS', attHome.Last10StreackWinning , attAway.Last10StreackWinning],
    ['INVICTO', attHome.Last10StreackNoLosing , attAway.Last10StreackNoLosing]
  ]);
  var chart = new google.charts.Bar(document.getElementById('racha'));

  chart.draw(ganados, google.charts.Bar.convertOptions(options));
}

/*-------------------------Grafica Probabilidades--------------------------- */
function tablaPorcentaje() {
  var posibilidades = google.visualization.arrayToDataTable([
    ['%', 'Probabilidades'],
    ['LOCAL', (pro1+pro2+proX)/pro1],
    ['EMPATE', (pro1+pro2+proX)/proX],
    ['VISITANTE', (pro1+pro2+proX)/pro2]
  ]);
  var chart = new google.visualization.PieChart(document.getElementById('porcentaje'));
  chart.draw(posibilidades, options2);
}






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