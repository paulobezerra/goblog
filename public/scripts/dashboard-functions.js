function confirmarExclusao(event) {
    if (!window.confirm('Você deseja confirmar exclusão?')) {
        event.preventDefault();
    }
}

function generateSlug(input, idInputSlug) {
    const title = input.value

    const slug = title.toLowerCase().replace(/ /g, '-').replace(/[^\w-]+/g, '');

    const inputSlug = document.getElementById(idInputSlug)

    if (inputSlug.value != slug) {
        inputSlug.value = slug
    }

}