$("#form-enroll-user").on("submit", enroll);

function enroll(e) {
    e.preventDefault();

    if ($('#password').val() != $('#repass').val()) {
        alert('Atention! Passwords no match');
        return;
    }

    $.ajax({
            url: '/users',
            method: 'POST',
            data: {
                name: $('#name').val(),
                nick: $('#nick').val(),
                email: $('#email').val(),
                password: $('#password').val(),
            },
        })
        .done(function () {
            alert('user enrolled success');
        })
        .fail(function (err) {
            alert(err);
        });
}