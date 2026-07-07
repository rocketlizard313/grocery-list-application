async function getData() {
  const url = "http://localhost:8080/grocery-list?id=1";
  try {
    const response = await fetch(url,{method:"GET"});
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const result = await response.json();
    console.log(result);
  } catch (error) {
    console.error(error.message);
  }

}

console.log("hello1");
getData();