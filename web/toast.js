export function toast(message, duration = 3000) {
    const toastContainer = document.getElementById('toast-container')

    const toast = document.createElement('div')
    toast.classList.add('toast')

    const toastMessage = document.createElement('div')
    toastMessage.classList.add('toast-message')
    toastMessage.textContent = message
    toast.appendChild(toastMessage)

    const toastClose = document.createElement('button')
    toastClose.classList.add('toast-close')
    toastClose.textContent = 'X'
    toastClose.addEventListener('click', () => toast.remove())
    toast.appendChild(toastClose)

    toastContainer.appendChild(toast)
    setTimeout(() => toast.remove(), duration)
}
