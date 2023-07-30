

let dataBlog = []

function addBlog(event) {

     event.preventDefault()

let name = document.getElementById("input-name").value
let start = new Date(document.getElementById("input-start").value)
let end = new Date(document.getElementById("input-end").value)
let description = document.getElementById("input-description").value
let image = document.getElementById("input-blog-image").files

const checkbox1 =  '<i class="fab fa-node"></i>'
const checkbox2 =  '<i class="fab fa-node-js"></i>'
const checkbox3 =  '<i class="fab fa-java"></i>'
const checkbox4 =  '<i class="fab fa-react"></i>'

let iconcheckbox1  = document.getElementById("checkbox1").checked ? checkbox1 : ""
let iconcheckbox2  = document.getElementById("checkbox2").checked ? checkbox2 : ""
let iconcheckbox3  = document.getElementById("checkbox3").checked ? checkbox3 : ""
let iconcheckbox4  = document.getElementById("checkbox4").checked ? checkbox4 : ""

image = URL.createObjectURL(image[0])

// 1000 ms = 1 detik
let milisecond = end - start

// 10000000 ms / 1000 = 10000 detik
let Second = Math.floor(milisecond / 1000)
// 10000 detik / 60 = 166 menit , bulatkan ke bawah
let Minutes = Math.floor(Second / 60)
// 166 menit / 60 = 2 jam 
let Hours = Math.floor(Minutes / 60)
// misalny 25 jam maka dia itung 1 hari dan kurg 24 jam = 0 hari
let Days = Math.floor(Hours / 24)
// konsepny sama persis
let Weeks = Math.floor(Days / 7)
let Months = Math.floor(Weeks / 4)
let Years = Math.floor(Months / 12)


if (Days < 1) {
    keterangan = "Waktu Telah Habis"
} else if (Days > 0 && Days < 7) {
   keterangan  = Days + "hari";
} else if (Weeks < 4) {
   keterangan  = Weeks + "minggu";
} else if (Months < 12) {
   keterangan  = Months + "bulan";
} else  {
   keterangan  = Years + "tahun";
}





let blog = {
    name,
    keterangan,
    description,
    iconcheckbox1,
    iconcheckbox2,
    iconcheckbox3,
    iconcheckbox4,
    image
}


dataBlog.push(blog)

tampilinBlog()
console.log(dataBlog)

}

function tampilinBlog() {
 document.getElementById("myproject").innerHTML = ""



 for (let index = 0; index < dataBlog.length; index++) {
    document.getElementById("myproject").innerHTML += 
    `
    <div class="grid-item">
    <a href="detail_project.html">
      <img class="w-100" src="${dataBlog[index].image}" height="100px">
      <div style="text-align: start;">
        <h3 style="word-wrap: break-word;">${dataBlog[index].name}</h3><br>
        <p>${dataBlog[index].keterangan}</p><br>
        <p>Description</p>
        <span class="d-block" style="width: 300px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis; font-size: 5pt;">
            ${dataBlog[index].description}
        </span>
        <br>
        <p>
            ${dataBlog[index].iconcheckbox1}
            ${dataBlog[index].iconcheckbox2}
            ${dataBlog[index].iconcheckbox3}
            ${dataBlog[index].iconcheckbox4}
        </p>
        <div>
          <a><button class="btn btn-primary" style="width: 100px; margin-left: 20px;">Edit</button></a>
          <a><button class="btn btn-primary" style="width: 100px; margin-left: 20px;">Delete</button></a>
        </div>
      </div>
    </a>
  </div>
    `   
    
 }
 
}