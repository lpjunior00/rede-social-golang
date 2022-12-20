$("#form-newpost").on('submit', createPost)

function createPost(event){
    event.preventDefault();

    let title = $("#title").val();
    let content = $("#content").val();

    if (title == "" || content == ""){
        Swal.fire('Atenção!', 'É obrigatório informar o título e o conteúdo da publicação.', 'warning')
        return
    }

    $.ajax({
        url: "/posts",
        method: "post",
        data: {
            title: title,
            content: content
        }
    }).done(function (){
        Swal.fire('Sucesso!', 'Post criado com sucesso!', 'success')
        .then(function (){
            window.location = "/home";
        })
    }).fail(function (){
        Swal.fire('Oops!', 'Ocorreu um erro ao efetuar o login automatico.', 'error')
    })

}
