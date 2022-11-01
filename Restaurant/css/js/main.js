let seciones = document.querySelectorAll("section");
let navLinks = document.querySelectorAll("nav ul li a");
window.addEventListener("scroll", () => {
    seciones.forEach((section) => {
        let top = window.scrollY,
            offset = section.offsetTop,
            height = section.offsetHeight,
            id = section.getAttribute("id");
        if (top >= offset && top < offset + height) {
            navLinks.forEach((link) => {
                link.classList.remove("active");
                document
                    .querySelector("nav ul li a[href *=" + id + "]")
                    .classList.add("active");
            });
        }
    });
});

// in mobil media open and close menue
openMenue = document.getElementById("open-nav");
closeMenue = document.getElementById("close-nav");
menue = document.querySelector(".nav-links");
openMenue.addEventListener("click", () => {
    menue.classList.add("clopssing");
    openMenue.style.cssText = "opacity:0;";
});
document.addEventListener("click", (element) => {
    if (
        element.target.id !== "open-nav" &&
        !element.target.classList.contains("nav-links") &&
        element.target.parentElement.id !== "open-nav"
    ) {
        menue.classList.remove("clopssing");
        openMenue.style.cssText = "opacity:1;";
    }
});
// menue filter
let selectedLi = document.querySelectorAll(".selection li"),
    productContainer = document.querySelectorAll(".content .box");
selectedLi.forEach((li) => {
    li.addEventListener("click", (event) => {
        toggleActiveClass(event);
    });
    li.addEventListener("click", manageProduct);
});
// function
function toggleActiveClass(list) {
    // remove active class
    list.target.parentElement.parentElement
        .querySelectorAll(".active")
        .forEach((element) => {
            element.classList.remove("active");
        });
    // add active class
    list.target.classList.add("active");
}
function manageProduct() {
    productContainer.forEach((product) => {
        product.style.display = "none";
    });
    document.querySelectorAll(this.dataset.categroy).forEach((element) => {
        element.style.display = "block";
    });
}
// testimoanial
let slide = document.querySelectorAll(".card");
let index = 0;
function next() {
    slide[index].classList.remove("active");
    index = (index + 1) % slide.length;
    slide[index].classList.add("active");
}
function previous() {
    slide[index].classList.remove("active");
    index = (index - 1 + slide.length) % slide.length;
    slide[index].classList.add("active");
}
