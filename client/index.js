async function getData() {
  const url = "http://localhost:8080/grocery-list?id=2";
  try {
    const response = await fetch(url,{method:"GET"});
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const result = await response.json();



    console.log(result);


    let card = document.querySelector(".card > ul");


    if(!card )
    {
      console.log("No card found");
      return;
    }

    console.log(result.id);

    result.groceryItems.forEach((element)=>{
      let itemList = document.createElement("li");

      console.log(element);

      itemList.textContent = element;

      card.appendChild(itemList);
    })
  } catch (error) {
    console.error(error.message);
  }

}



console.log("hello1");
getData();

