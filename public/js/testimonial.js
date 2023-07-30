const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest()

  xhr.open("GET", "http://127.0.0.1:5000/data", true)
  xhr.onload = function () {
      // http code : 200 -> OK
      if (xhr.status === 200) {
          resolve(JSON.parse(xhr.responseText))
      } else if (xhr.status >= 400) {
          reject("Error loading data")
      }
  }
  xhr.onerror = function () {
      reject("Network error")
  }
  xhr.send()
})


let testimonialData = []

async function getData(rating) {
  try {
      const response = await promise
      console.log(response)
      testimonialData = response
      allTestimonial()
  } catch (err) {
      console.log(err)
  }
}

getData()

function allTestimonial() {
  let testimonialHTML = ""

  testimonialData.forEach((card) => {
      testimonialHTML += `
      
      <div class="mt-4 col-md-4 ">
      <div class="bg-white">
        <img src="${card.image}" class="w-100 border mt-2" alt="" style="height: 100px;">
        <h3>${card.quote}</h3>
        <p class="text-right">${card.user}
        <span></span>
        </p>
        <p class="author text-right">${card.rating} <i class="fa-solid fa-star"></i></p>
      </div>
    </div>

`
  })

  document.getElementById("testimonials").innerHTML = testimonialHTML
}

function filterTestimonial(rating) {
  let filteredTestimonialHTML = ""

  const filteredData = testimonialData.filter((card) => {
      return card.rating === rating
  }) 

  filteredData.forEach((card) => {
      filteredTestimonialHTML += `
      <div class="mt-4 col-md-4 ">
      <div class="bg-white">
        <img src="${card.image}" class="w-100 border mt-2" alt="" style="height: 100px;">
        <h3>${card.quote}</h3>
        <p class="text-right">${card.user}</p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
      </div>
    </div>  
  `
  })

  document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}



