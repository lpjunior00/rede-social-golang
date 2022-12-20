$("#btn-follow").on("click", follow);
$("#btn-unfollow").on("click", unfollow);
$("#btn-editUser").on("click", editUser);
$("#btn-editPassword").on("click", editPassword);
$("#btn-return").on("click", returnToProfile)
$("#btn-delete-account").on("click", deleteAccount)


function follow(){
    const userId = $(this).data('user-id')
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${userId}/follow`,
        method: "POST"
    }).done(function (){
        window.location = `/users/${userId}`
    }).fail(function (){
        Swal.fire("oops", "erro ao seguir usuário", "error")
        $("#btn-follow").prop('disabled', false)  
    })
}

function unfollow(){
    const userId = $(this).data('user-id')
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: "POST"
    }).done(function (){
        window.location = `/users/${userId}`
    }).fail(function (){
        Swal.fire("oops", "erro ao parar de seguir usuário", "error")
        $("#btn-unfollow").prop('disabled', false)  
    })
}

function editUser(event){
    event.preventDefault();
    
    $(this).prop('disabled', true)
    $.ajax({
        url: '/edit-user',
        method: 'PUT',
        data: {
            name: $("#name").val(),
            email: $("#email").val(),
            nickname: $("#nickname").val(),
        }
    }).done(function (){
        Swal.fire("Sucesso!", "Dados atualizados com sucesso!", "success")
            .then(function () {
                window.location = "/profile"
            })
    }).fail(function () {
        Swal.fire("Oops", "Erro ao atualizar dados do usuário!", "error")
    })

}

function editPassword(event){
    event.preventDefault();

    const newPassword = $("#newPassword").val();
    const confirmNewPassword = $("#confirmNewPassword").val();
    if (newPassword != confirmNewPassword){
        Swal.fire("Oops", "A nova senha não coincide com o confirmar nova senha.", "error")
        return
    }
    
    $(this).prop('disabled', true)
    $.ajax({
        url: '/change-password',
        method: 'POST',
        data: {
            currentPassword: $("#currentPassword").val(),
            newPassword: newPassword,
        }
    }).done(function (){
        Swal.fire("Sucesso!", "Senha atualizada com sucesso!", "success")
            .then(function () {
                window.location = "/profile"
            })
    }).fail(function () {
        Swal.fire("Oops", "Erro ao atualizar a senha do usuário!", "error")
        $("#btn-editPassword").prop("disabled", false)
    })

}

function returnToProfile(){
    window.location = "/profile";
}

function deleteAccount(){
    Swal.fire({
        title:"Atenção",
        text: "Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversivel.",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "Warning"
    }).then(function(confirmacao){
        if (confirmacao.value){
            $.ajax({
                url: "/users",
                method: "DELETE"
            }).done(function (){
                Swal.fire("Sucesso!", "Conta permanentemente excluida do sistema.", "success")
                    .then(function (){
                        window.location = "/logout"
                    })
            }).fail(function (){
                Swal.fire("Ops!", "Ocorreu um erro ao excluir a conta", "error")
            })
        }
    })
}