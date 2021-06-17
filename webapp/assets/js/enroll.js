$('#form-enroll-user').on('submit', enroll);

function enroll(e) {
	e.preventDefault();

	if ($('#password').val() != $('#repass').val()) {
		swal('Atention!', 'Passwords no match', 'error');
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
		.done(() => {
			swal('Sucesso!', 'User created!', 'success').then(() => {
				$.ajax({
					url: '/login',
					method: 'POST',
					data: {
						email: $('#email').val(),
						password: $('#password').val(),
					},
				}).done(() => {
					window.location = '/home';
				});
			});
		})
		.fail((err) => {
			swal('Atention!', err, 'error');
		});
}
