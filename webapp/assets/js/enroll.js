$('#form-enroll-user').on('submit', enroll);

function enroll(e) {
	e.preventDefault();

	if ($('#password').val() != $('#repass').val()) {
		toastr.options = { "progressBar": true };
		toastr.error('Passwords no match');
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
			$.ajax({
				url: '/login',
				method: 'POST',
				data: {
					email: $('#email').val(),
					password: $('#password').val(),
				},
			}).done(() => {
				toastr.success('User was created');
				window.location = '/home';
			});
		})
		.fail(() => {
			toastr.error('Error occurred when creating user');
		});
}
