function asyncFunction() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve('Async Hello World')
    }, 16);
  })
}

asyncFunction().then((value) => {
  console.log(value)
}).catch((error) => {
  console.log(error)
})
