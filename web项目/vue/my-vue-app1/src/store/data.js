export default Array(4000).fill(1).map((el, index) => {
    return {
        id: index,
        content: '',
        width: null,
        select: false,
        color: ''
    }
}).map((el, index) => {
    let y = Math.floor((index + 1) / 50) + 1
    let x = (index + 1) % 50
    el.x = x
    el.y = y
    return el
})