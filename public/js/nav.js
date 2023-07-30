var navContainer = document.getElementById("navContainer");

var navHTML = `
    <nav class="navbar navbar-expand-lg navbar-light bg-light ">
        <a class="navbar-brand" href="#">
            <img src="../public/image/logo.png" width="50" height="30" alt="Logo">
        </a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
            aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse text-center" id="navbarNav">
            <ul class="navbar-nav">
                <li class="nav-item active">
                    <a class="nav-link" href="/">Home</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link active test2" href="myproject">My Project</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link active test2" href="testimonial">Testimonial</a>
                </li>
            </ul>
            <div class="nav-right ml-auto">
                <a class="nav-link btn rounded-pill contact-button" href="contact" style="text-decoration: none; background-color:orange;  color:black;">Contact Me</a>
            </div>
        </div>
    </nav>
`;

navContainer.innerHTML = navHTML;







// Add an event listener to the toggle button
var toggleButton = document.querySelector(".navbar-toggler");
var contactButton = document.querySelector(".contact-button");
var home = document.querySelector(".test");
toggleButton.addEventListener("click", function() {
    if (window.innerWidth <= 990 && toggleButton.getAttribute("aria-expanded") === "true") {
        contactButton.style.backgroundColor = "transparent";
        contactButton.style.borderRadius = "-1";

        home.style.textAlign = "center";
    } else {
        contactButton.style.backgroundColor = "";
        contactButton.style.borderRadius = "";
        home.style.textAlign = "";
    }
});

// Add a resize event listener to handle changes in screen width
window.addEventListener("resize", function() {
    if (window.innerWidth <= 990 && toggleButton.getAttribute("aria-expanded") === "true") {
        contactButton.style.backgroundColor = "transparent";
        contactButton.style.borderRadius = "-1";
        home.style.textAlign = "center";
    } else {
        contactButton.style.backgroundColor = "";
        contactButton.style.borderRadius = "";

    }
});

