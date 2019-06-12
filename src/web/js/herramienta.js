/*EVENTOS CARGA DOCUMENTO*/

var arr

$(document).ready(function () {
  path = "http://localhost:3010/api/v1/"

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

  $(function () {
    recibirPaises();
  });

  $("select[name='Paises']").change(function () {
    recibirDivisiones();
  });

  $("select[name='Division']").change(function () {
    recibirLocales();

  });

  $("select[name='Local']").change(function () {
    recibirVisitantes();
  });

  $("select[name='Visitante']").change(function () {
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
  var pais = "" + $('select[name="Paises"] option:selected').text();
  $.ajax({
    url: path + pais + "/divisions",
    success: function (result) {
      insertarDivisiones(result);
    }
  });
};

/*solicito equipos al servidor*/
function recibirLocales() {
  var pais = $('select[name="Paises"] option:selected').text(),
    division = $('select[name="Division"] option:selected').text();
  console.log(pais, division)
  $.ajax({
    url: path + pais + "/" + division + "/teams",
    success: function (result) {
      insertarLocales(result);
    }
  });
};

/*solicito equipos al servidor*/
function recibirVisitantes() {
  let pais = $('select[name="Paises"] option:selected').text(),
    division = $('select[name="Division"] option:selected').text();
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
  for (let i = 0; i < p.length; i++) {
    let option = $("<option></option>");
    $(option).html(p[i]);
    $("select[name='Paises']").append(option);
  }
};

/*inserto divisiones en el select divisiones*/
function insertarDivisiones(p) {
  $("select[name='Division']").empty();
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
  for (let i = 0; i < p.length; i++) {
    let option = $("<option></option>");
    $(option).html(p[i]);
    $("select[name='Visitante']").append(option);
  }

};








/*----------------------- Atributos en pagina--------------------------*/


//path+pais + "/" + division + "/" + local + "/properties",
function solicitarAtributos() {
  var pais = "" + $('select[name="Paises"] option:selected').text();
  var division = "" + $('select[name="Division"] option:selected').text();
  var local = "" + $('select[name="Local"] option:selected').text();
  var visitante = "" + $('select[name="Visitante"] option:selected').text();
  arr = []
  $.ajax({
    url: "http://localhost:3010/api/v1/Spain/SP1/Girona/properties",
    success: function (home) {
      console.log(home)
      $.ajax({
        url: "http://localhost:3010/api/v1/Spain/SP1/Girona/properties",
        success: function (away) {
          console.log(away)
          arr["home"] = home
          arr["away"] = away
          console.log(arr.home)
          insertarAtributos()
        }
      });
    }
  });
};

function insertarAtributos() {
  $("#gLocal").text(arr.home.Last10WinningLocalMatchs);
  $("#eLocal").text(arr.home.Last10TiedingLocalMatchs);
  $("#pLocal").text(arr.home.Last10LosingLocalMatchs);
  $("#gVisitante").text(arr.away.Last10WinningAwayMatchs);
  $("#eVisitante").text(arr.away.Last10TiedingAwayMatchs);
  $("#pVisitante").text(arr.away.Last10LosingAwayMatchs);
  $("#gmLocal").text(arr.home.Last10GoalsTuckedAmountLocalMatchs);
  $("#gmVisitante").text(arr.away.Last10GoalsTuckedAmountLocalMatchs);
  $("#geLocal").text(arr.home.Last10GoalsReceivedAmountLocalMatchs);
  $("#geVisitante").text(arr.away.Last10GoalsReceivedAmountLocalMatchs);
  $("#rLocal").text(arr.home.Last10StreackWinningLocal);
  $("#rVisitante").text(arr.away.Last10StreackWinningAway);
  $("#iLocal").text(arr.home.Last10LosingLocalMatchs);
  $("#iVisitante").text(arr.away.Last10StreackNoLosingLocal);
  pro1 =0.027 * arr.home.Last10WinningLocalMatchs +
    -0.0134 * arr.away.Last10WinningAwayMatchs +
    -0.0303 * arr.away.Last10TiedingAwayMatchs+
     0.114 * arr.home.Last10GoalsTuckedAmountLocalMatchs +
    -0.0483 * arr.home.Last10GoalsReceivedAmountLocalMatchsl +
    -0.0882 * arr.away.Last10GoalsReceivedAmountLocalMatchs +
    -0.0604 * arr.home.Last10AverageTuckedGoalsLocal +
    0.0888 * arr.home.Last10AverageReceivedGoalsLocal +
    0.4706;
  pro2 = -0.0132 * arr.home.Last10TiedingLocalMatchs +
    0.0358 * arr.away.Last10WinningAwayMatchs +
    -0.0215 * arr.home.Last10GoalsTuckedAmountLocalMatchs +
    0.2181 * arr.home.Last10GoalsReceivedAmountLocalMatchs +
    -0.0433 * arr.away.Last10GoalsTuckedAmountLocalMatchs +
    0.0226 * arr.home.Last10AverageTuckedGoalsLocal +
    -0.1045 * arr.home.Last10AverageReceivedGoalsLocal +
    0.336;
  proX = 1 - (pro1 + pro2);
  $("#proLocal").text(pro1);
  $("#proEmpate").text(proX);
  $("#proVisitante").text(pro2);
  $("#localA").val(pro1);
  $("#emmpateA").val(proX);
  $("#visitanteA").val(pro2);
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
    ["GANADOS", arr.home.Last10WinningLocalMatchs, arr.away.Last10WinningAwayMatchs],
    ["EMPATADOS", arr.home.Last10TiedingLocalMatchs, arr.away.Last10TiedingAwayMatchs],
    ["PERDIDOS", arr.home.Last10LosingLocalMatchs,arr.away.Last10LosingAwayMatchs]
  ]);
  var chart = new google.charts.Bar(document.getElementById('ultimos'));

  chart.draw(resultados, google.charts.Bar.convertOptions(options));
}


/*------------------Grafica GOLES-----------------------------*/
function tablaGoles() {
  var goles = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['A FAVOR', arr.home.Last10GoalsTuckedAmountLocalMatchs, arr.away.Last10GoalsTuckedAmountLocalMatchs],
    ['EN CONTRA', arr.home.Last10GoalsReceivedAmountLocalMatchs, arr.away.Last10GoalsReceivedAmountLocalMatchs]
  ]);
  var chart = new google.charts.Bar(document.getElementById('goles'));

  chart.draw(goles, google.charts.Bar.convertOptions(options));
}

/*------------------Grafica RACHA-------------------------------------*/
function tablaRachaG() {
  var ganados = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['GANADOS', arr.home.Last10StreackWinningLocal , arr.away.Last10StreackWinningAway ],
    ['INVICTO', arr.home.Last10StreackNoLosingLocal , arr.away.Last10StreackNoLosingLocal ]
  ]);
  var chart = new google.charts.Bar(document.getElementById('racha'));

  chart.draw(ganados, google.charts.Bar.convertOptions(options));
}

/*-------------------------Grafica Probabilidades--------------------------- */
function tablaPorcentaje() {
  var posibilidades = google.visualization.arrayToDataTable([
    ['%', 'Probabilidades'],
    ['LOCAL', pro1],
    ['EMPATE', proX],
    ['VISITANTE', pro2]
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