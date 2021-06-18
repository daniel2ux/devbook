$('#new-post').on('submit', newPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.dislike-post', dislikePost);
$('#btn-update').on('click', updatePost);
$('.delete-post').on('click', deletePost);

function newPost(e) {
	e.preventDefault();

	$.ajax({
		url: '/posts',
		method: 'POST',
		data: {
			title: $('#title').val(),
			post: $('#post').val(),
		},
	})
		.done(() => {
			window.location = '/home';
		})
		.fail(() => {
			alert('Publish fail');
		});
}

function likePost(e) {
	e.preventDefault();
	const clickedElement = $(e.target);
	const postID = clickedElement.closest('.card').data('post-id');
	clickedElement.prop('disabled', true);

	$.ajax({
		url: `posts/${postID}/like`,
		method: 'POST',
	})
		.done(() => {
			const likeCounter = clickedElement.closest('.fa-heart');
			const likeQtdy = parseInt(likeCounter.text());
			likeCounter.text(likeQtdy + 1);

			clickedElement.addClass('dislike-post');
			clickedElement.addClass('text-danger');
			clickedElement.removeClass('like-post');
		})
		.fail(() => {
			alert('Liked fail!');
		})
		.always(() => {
			clickedElement.prop('disabled', false);
		});
}

function dislikePost(e) {
	e.preventDefault();
	const clickedElement = $(e.target);
	const postID = clickedElement.closest('.card').data('post-id');
	clickedElement.prop('disabled', true);

	$.ajax({
		url: `posts/${postID}/dislike`,
		method: 'POST',
	})
		.done(() => {
			const likeCounter = clickedElement.next('span');
			const likeQtdy = parseInt(likeCounter.text());
			likeCounter.text(likeQtdy - 1);

			clickedElement.removeClass('dislike-post');
			clickedElement.removeClass('text-danger');
			clickedElement.addClass('like-post');
		})
		.fail(() => {
			alert('dislike fail!');
		})
		.always(() => {
			clickedElement.prop('disabled', false);
		});
}

function updatePost() {
	$(this).prop('disabled', true);

	const postID = $(this).data('post-id');

	$.ajax({
		url: `/posts/${postID}`,
		method: 'PUT',
		data: {
			title: $('#title').val(),
			post: $('#post').val(),
		},
	})
		.done(() => {
			swal({
				title: 'Devbook',
				text: 'update success...',
				icon: 'success',
				buttons: false,
				timer: 2000,
			}).then(function () {
				window.location = '/home';
			});
		})
		.fail((err) => {
			swal({
				title: 'Devbook',
				text: 'update fail!',
				icon: 'error',
			});
		})
		.always(() => {
			$('#btn-update').prop('disabled', false);
		});
}

function deletePost(e) {
	e.preventDefault();
	const clickedElement = $(e.target);
	const post = clickedElement.closest('.card');
	const postID = post.data('post-id');
	clickedElement.prop('disabled', true);

	swal({
		title: 'Atention!',
		text: 'The publication will be permanently deleted. Are you sure?',
		icon: 'warning',
		buttons: true,
		dangerMode: true,
	}).then((willDelete) => {
		if (willDelete) {
			$.ajax({
				url: `/posts/${postID}`,
				method: 'DELETE',
			})
				.done(() => {
					post.fadeOut('slow', function () {
						$(this).remove();
					});
				})
				.fail((err) => {
					alert('delete fail!');
				})
				.always(() => {
					clickedElement.prop('disabled', false);
				});
		}
	});
}
