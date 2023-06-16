import { ref, reactive } from 'vue'
import dataRow from "./data";

const data = ref(dataRow)
const activeline = reactive({ x: '', y: '' })

const copy = (id) => {

}
const setColor = (id, color) => {
    data.value[id].color = color
}
const selectLine = (item, xOry) => {
    activeline.x = ''
    activeline.y = ''
    activeline[xOry] = item
    console.log(item.xOry);
    data.value.forEach(el => {
        el.select = false
        if (el[xOry] == item) {
            el.select = true
        }
    })
}
export default { data, activeline, copy, setColor, selectLine }