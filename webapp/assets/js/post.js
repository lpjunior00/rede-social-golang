$(document).on('click', ".like-post", likePost);
$(document).on('click', ".dislike-post", dislikePost);

$('#btn-edit').on('click', updatePost);
$('#btn-delete').on('click', deletePost);

function likePost(event){
    event.preventDefault();

    const elemento = $(event.target);
    const postId = elemento.closest('div').data('post-id')
    
    $.ajax({
        url: `/posts/${postId}/like`,
        method: "POST"
    }).done(function () {
        const textLikes = elemento.next('span')
        const nuLikes = parseInt(textLikes.text());

        textLikes.text(nuLikes + 1)
        elemento.addClass('dislike-post');
        elemento.addClass('text-danger');
        elemento.removeClass('like-post');
        
    }).fail(function () {
        console.log("fail")
    })

}

function dislikePost(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const postId = elementoClicado.closest('div').data('post-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/dislike`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('dislike-post');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('like-publicacao');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function updatePost() {
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');
    
    $.ajax({
        url: `/posts/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        Swal.fire('Sucesso!', 'Publicação atualizada com sucesso!', 'success')
            .then(function() {
                window.location = "/home";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    })
}

function deletePost(evento) {
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoId = publicacao.data('post-id');
    
        elementoClicado.prop('disabled', true);
    
        $.ajax({
            url: `/posts/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });
    })

}