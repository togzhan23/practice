const links = document.querySelectorAll('.nav-link');
const pages = document.querySelectorAll('.page');

links.forEach(link => {
  link.addEventListener('click', e => {
    e.preventDefault();

    const target = link.getAttribute('data-page');

    // Переключение активной ссылки
    links.forEach(l => l.classList.remove('active'));
    link.classList.add('active');

    // Переключение активной страницы
    pages.forEach(page => {
      if (page.id === target) {
        page.classList.add('active');
      } else {
        page.classList.remove('active');
      }
    });
  });
});
