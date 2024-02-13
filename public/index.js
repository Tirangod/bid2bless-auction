console.log("index.js: started!")

import Alpine from 'resourses/js/alpinejs/alpinejs.js'

async function loadHTMLComponent(name) {
    // TODO: handle errors
    let html = ""
    await fetch("components/" + name + ".html").
        then(value => value.text()).catch(err => console.error(err))
        then(text => html = text).catch(err => console.error(err))
    return html
}

Alpine.directive("comp", async (el, { expression }, { evaluate }) => {
    const compName = expression
    const html = await loadHTMLComponent(compName)
    el.innerHTML = html
})

window.Alpine = Alpine

Alpine.start()