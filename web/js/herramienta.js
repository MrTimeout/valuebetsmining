/*EVENTOS CARGA DOCUMENTO*/

$(document).ready(function () {


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
    insertarAtributos();
  });

});






/*---------------FUNCIONES  DE SELECCION---------------------------*/

/*inicio variables globales*/
var paises = [];
var divisiones = [];
var locales = [];
var visitantes = [];
var atributos;



/*solicito paises al servidor*/
function recibirPaises() {
  $.ajax({
    url: "api/v1/countries",
    success: function (result) {
      paises = [];
      for (let i = 0; i < result.length; i++) {
        paises.push(result[i]);
      }
    }
  });

  insertarPaises();
};

/*solicito divisiones al servidor*/
function recibirDivisiones() {
  var pais = "" + $('select[name="Paises"] option:selected').text();
  $.ajax({
    url: "/api/v1/" + pais,
    success: function (result) {
      divisiones = [];
      for (let i = 0; i < result.length; i++) {
        divisiones.push(result[i]);
      }
    }
  });
  insertarDivisiones();
};

/*solicito equipos al servidor*/
function recibirLocales() {
  var pais = "" + $('select[name="Paises"] option:selected').text();
  var division = +$('select[name="Division"] option:selected').text();
  $.ajax({
    url: "/api/v1/" + pais + "/" + division,
    success: function (result) {
      locales = [];
      for (let i = 0; i < result.length; i++) {
        locales.push(result[i]);
      }
    }
  });


  insertarLocales();
};

/*solicito equipos al servidor*/
function recibirVisitantes() {
  var pais = "" + $('select[name="Paises"] option:selected').text();
  var division = +$('select[name="Division"] option:selected').text();
  $.ajax({
    url: "/api/v1/" + pais + "/" + division,
    success: function (result) {
      visitantes = [];
      for (let i = 0; i < result.length; i++) {
        visitantes.push(result[i]);
      }
    }
  });

  insertarVisitantes();
};

/*inserto paises en el select paises*/
function insertarPaises() {
  $("select[name='Paises']").empty();
  for (var i = 0; i < paises.length; i++) {
    var option = $("<option></option>");
    $(option).html(paises[i]);
    $("select[name='Paises']").append(option);
  }
};

/*inserto divisiones en el select divisiones*/
function insertarDivisiones() {
  $("select[name='Division']").empty();
  for (var i = 0; i < divisiones.length; i++) {
    var option = $("<option></option>");
    $(option).html(divisiones[i]);
    $("select[name='Division']").append(option);
  }
  $("select[name='Local']").empty();
  $("select[name='Visitante']").empty();
};

/*inserto equipos en el select locales*/
function insertarLocales() {
  $("select[name='Local']").empty();
  for (var i = 0; i < locales.length; i++) {
    var option = $("<option></option>");
    $(option).html(locales[i]);
    $("select[name='Local']").append(option);
  }
  $("select[name='Visitante']").empty();
};

/*inserto equipos en el select visitantes*/
function insertarVisitantes() {
  $("select[name='Visitante']").empty();
  for (var i = 0; i < visitantes.length; i++) {
    var option = $("<option></option>");
    $(option).html(visitantes[i]);
    $("select[name='Visitante']").append(option);
  }

};








/*----------------------- Atributos en pagina--------------------------*/

function solicitarAtributos() {
  var pais = "" + $('select[name="Paises"] option:selected').text();
  var division = "" + $('select[name="Division"] option:selected').text();
  var local = "" + $('select[name="Local"] option:selected').text();
  var visitante = "" + $('select[name="Visitante"] option:selected').text();
  $.ajax({
    url: "/api/v1/" + pais + "/" + division + "/" + local + "/" + visitante,
    success: function (result) {


      console.log(result);

      atributos = result;

    }
  });
};

function insertarAtributos() {
  $("#gLocal").text();
  $("#eLocal").text();
  $("#pLocal").text();
  $("#gVisitante").text();
  $("#eVisitante").text();
  $("#pVisitante").text();
  $("#gmLocal").text();
  $("#gmVisitante").text();
  $("#geLocal").text();
  $("#geVisitante").text();
  $("#rLocal").text();
  $("#rVisitante").text();
  $("#iLocal").text();
  $("#iVisitante").text();
  var pro1 =
    0.027 * ganadosLocal +
    -0.0134 * ganadosVisitante +
    -0.0303 * empatadosVisitante +
    0.114 * marcadosLocal +
    -0.0483 * recibidosLocal +
    -0.0882 * recibidosvisitante +
    -0.0604 * mediaMarcadosLocal +
    0.0888 * mediaEncajadosLocal +
    0.4706;
  var pro2 = -0.0132 * empatadosLocal +
    0.0358 * ganadosVisitante +
    -0.0215 * marcadosLocal +
    0.2181 * recibidosLocal +
    -0.0433 * marcadosvisitante +
    0.0226 * mediaMarcadosLocal +
    -0.1045 * mediaEncajadosLocal +
    0.336;
  var proX = 1 - (pro1 + pro2);
  $("#proLocal").text(pro1);
  $("#proEmpate").text(proX);
  $("#proVisitante").text(pro2);
  $("#localA").val(pro1);
  $("#emmpateA").val(prox);
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



/*--------------Grafica RESULTADOS-------------------------*/
function drawChart() {
  var resultados = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['GANADOS', 2, 2],
    ['EMPATADOS', 2, 2],
    ['PERDIDOS', 3, 6]
  ]);
  var chart = new google.charts.Bar(document.getElementById('ultimos'));

  chart.draw(resultados, google.charts.Bar.convertOptions(options));
}


/*------------------Grafica GOLES-----------------------------*/
function tablaGoles() {
  var goles = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['A FAVOR', 5, 2],
    ['EN CONTRA', 8, 2]
  ]);
  var chart = new google.charts.Bar(document.getElementById('goles'));

  chart.draw(goles, google.charts.Bar.convertOptions(options));
}

/*------------------Grafica RACHA-------------------------------------*/
function tablaRachaG() {
  var ganados = google.visualization.arrayToDataTable([
    ['Nº', 'Local', 'Visitante'],
    ['GANADOS', 5, 4],
    ['INVICTO', 8, 1]
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