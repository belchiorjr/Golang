
$("#nova-publicacao").on('submit', criarPublicacao)
$("#editar-publicacao").on('submit', salvarPublicacao)

$(document).on('click', '.curtir-publicacao', curtirPublicacao) 
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao) 
$(document).on('click', '.to-delete', excluirPublicacao) 


function criarPublicacao(evento) {
  evento.preventDefault()

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    },
    complete: function (response) {
      if (response.status == 200) {
        Swal.fire("Sucesso! ", "A publicação foi criada!", "success").then(function () {
          location.href= "/home"
        })
      } else {
        Swal.fire("Erro! ", "Erro ao criar a publicação.", "error")
      }
    }
  })
}


function salvarPublicacao(evento) {
  evento.preventDefault()
  var elementoClicado = $(evento.target)
  var publicacaoID = elementoClicado.closest('form').data("publicacao-id")

  elementoClicado.prop('disabled', true)

  $.ajax({
    url: `/publicacoes/${publicacaoID}`,
    method: "PUT",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val(),
    }
  }).done(function () {
    Swal.fire("Sucesso! ", "A publicação foi salva", "success").then(function () {
      location.href= "/home"
    })
  }).fail(function (erro) {
    Swal.fire("Erro! ", "Erro ao salvar a publicação", "error")
  }).always(function () {
    elementoClicado.prop('disabled', false)
  })
}


function excluirPublicacao(evento) {

  Swal.fire({
    title: "Atenção!",
    text: "Tem certeza que deseja excluir esta publicação? Essa ação é irreversível!",
    showCancelButton: true,
    cancelButtonText: "Cancelar",
    icon: "warning"
  }).then(function (confirmacao) {
    if (!confirmacao.value)
      return

    var elementoClicado = $(evento.target)
    var publicacaoID = elementoClicado.closest('.card').data("publicacao-id")

    elementoClicado.prop('disabled', true)

    $.ajax({
      url: `/publicacoes/${publicacaoID}`,
      method: "DELETE"
    }).done(function () {
      Swal.fire("Sucesso! ", "A publicação foi excluída", "success").then(function () {
        window.location = "/home"
      })
    }).fail(function (erro) {   
      Swal.fire("Erro! ", "Erro ao excluir a publicação", "error")
    }).always(function () {
      elementoClicado.prop('disabled', false)
    })
  })
}


function curtirPublicacao(evento) {
  var elementoClicado = $(evento.target)
  var publicacaoID = elementoClicado.closest('.card').data("publicacao-id")
  
  elementoClicado.prop('disabled', true)

  $.ajax({
    url: `/publicacoes/${publicacaoID}/curtir`,
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val(),
    }
  }).done(function () {
    const contadorDeCurtidas = elementoClicado.next('span')
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())
    contadorDeCurtidas.text(quantidadeDeCurtidas + 1)
    elementoClicado.addClass('descurtir-publicacao')
    elementoClicado.addClass('text-danger')
    elementoClicado.removeClass('curtir-publicacao')
    Swal.fire("Sucesso! ", "A publicação foi curtida", "success")
  }).fail(function (erro) {
    Swal.fire("Erro! ", "Não foi possível curtir a publicação", "error")
  }).always(function () {
    elementoClicado.prop('disabled', false)
  })
}

function descurtirPublicacao(evento) { 
  var elementoClicado = $(evento.target)
  var publicacaoID = elementoClicado.closest('.card').data("publicacao-id")
  
  elementoClicado.prop('disabled', true)

  $.ajax({
    url: `/publicacoes/${publicacaoID}/descurtir`,
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val(),
    }
  }).done(function () {
    const contadorDeCurtidas = elementoClicado.next('span')
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())
    contadorDeCurtidas.text(quantidadeDeCurtidas - 1)
    elementoClicado.addClass('curtir-publicacao')
    elementoClicado.removeClass('text-danger')
    elementoClicado.removeClass('descurtir-publicacao')
    Swal.fire("Sucesso! ", "A publicação foi descurtida", "success")
  }).fail(function (erro) {
    Swal.fire("Erro! ", "Não foi possível descurtir a publicação", "error")
  }).always(function () {
    elementoClicado.prop('disabled', false)
  })
}