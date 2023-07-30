var cardContainer = document.createElement("div");
cardContainer.className = "card container";

cardContainer.innerHTML = `
    <div class="card-body">
        <h1 class="card-title">Hi Welcome to My Blog</h1>
        <p class="card-text">Lorem ipsum dolor sit amet consectetur adipisicing elit. Nam, soluta cupiditate.
            Corrupti aliquid minima dicta quidem officia. Tempore officia unde impedit optio temporibus,
            quas eligendi dolorem ullam quasi asperiores voluptates.</p>
    </div>

    <div class="text-left">
        <a href="https://api.whatsapp.com/send/?phone=%2B6285715231512&text&type=phone_number&app_absent=0"
            target="_blank" class="btn" style="background-color:black; color:white ; border-radius:50px">Contact Me</a>
    </div>

    <div class="text-left">
        <a href="https://instagram.com/erwinz1412" target="_blank" class="btn">
            <img class="image1" src="../public/image/ig.jpeg" style="width: 20px; height: 20px;">
        </a>
        <a href="https://www.linkedin.com/in/erwin-00b580215/" target="_blank" class="btn">
            <img class="image1" src="../public/image/link.png" alt="" style="width: 20px; height: 20px;">
        </a>
        <a href="https://github.com/erwin1412" target="_blank" class="btn">
            <img class="image1" src="../public/image/git.png" alt="" style="width: 20px; height: 20px;">
        </a>
    </div>

    <div class="text-left">
        <br>
        <a href="https://drive.google.com/drive/folders/1VToHf6FXNYrkXiTPUxWbd3jWQ-bu8NTs?usp=sharing"
            target="_blank" class="btn">Download My CV <img src="../public/image/dwnld.jpg" alt=""
                style="width: 20px; height: 20px;"></a>
    </div>

    <div class="content1">
        <div class="text-center">
            <br>
            <img src="../public/image/profile.jpg" alt="" class="img-thumbnail" style="width: 250px; height: 250px;">
        </div>
        <div class="text-center">
            <h3>ERWIN</h3>
            <p>Fullstack Laravel</p>
        </div>
    </div>
`;

document.getElementById("cardContainer").appendChild(cardContainer);

