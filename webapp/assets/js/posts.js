$('#new-post').on('submit', newPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.dislike-post', dislikePost);

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
        .done(function () {
            window.location = '/home';
        })
        .fail(function () {
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
        .done(function () {
            const likeCounter = clickedElement.next('span');
            const likeQtdy = parseInt(likeCounter.text());
            likeCounter.text(likeQtdy + 1);

            clickedElement.addClass('dislike-post');
            clickedElement.addClass('text-danger');
            clickedElement.removeClass('like-post');
        })
        .fail(function () {
            alert('Liked fail!');
        })
        .always(function () {
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
        .done(function () {
            const likeCounter = clickedElement.next('span');
            const likeQtdy = parseInt(likeCounter.text());
            likeCounter.text(likeQtdy - 1);

            clickedElement.removeClass('dislike-post');
            clickedElement.removeClass('text-danger');
            clickedElement.addClass('like-post');
        })
        .fail(function () {
            alert('dislike fail!');
        })
        .always(function () {
            clickedElement.prop('disabled', false);
        });
}