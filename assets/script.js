const goWasm = new Go(); // Create a new Go instance
WebAssembly.instantiateStreaming(
  fetch("assets/main.wasm"),
  goWasm.importObject
).then((result) => {
    goWasm.run(result.instance); // Run the Go program
    document.getElementById("get-html").addEventListener("click", () => {
        document.body.innerHTML += getHtml()
    })});
