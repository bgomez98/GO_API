const getTasks = async () => {
  try {
    const response = await fetch('/api/todo-list/task')
    const data = await response.json() 
    return data
  } catch (err) {
    throw new Error(err.message)
  }
}

const renderTasks = (tasks) => {
  console.log(tasks)
  tasks.forEach((item) => {
  })
}

const renderButton = () => {
  const control = document.createElement('div')
  const button = document.createElement('button')
  
  control.setAttribute('class', 'control')
  button.setAttribute('class','button is-info')
  button.innerText = "Agregar"
  
  control.appendChild(button)

  return button
}

const renderInput = () => {
  const control = document.createElement('div')
  const input = document.createElement('input')

  control.setAttribute('class', 'control')
  input.setAttribute('class', 'input')
  input.setAttribute('type', 'text')
  input.setAttribute('placeholder', 'Ingresar nueva tarea')

  control.appendChild(input)

  return control
}

const renderMain = () => {
  const app = document.querySelector('#app') 
  const container = document.createElement('div')
  const box = document.createElement('div')
  const columns = document.createElement('div') 
  const column = document.createElement('div')
  const field = document.createElement('div')
  const button = renderButton()
  const input = renderInput()

  container.setAttribute('class', 'container my-5')
  box.setAttribute('class', 'box')
  columns.setAttribute('class', 'columns is-mobile is-centered')
  column.setAttribute('class', 'column is-10')
  column.setAttribute('id', 'tasks')
  field.setAttribute('class', 'field has-addons') 
  
  field.appendChild(input)
  field.appendChild(button)
  box.appendChild(field)
  column.appendChild(box)
  columns.appendChild(column)
  container.appendChild(columns)
  app.appendChild(container)

}

const init = async () => {
  const tasks = await getTasks()
  renderMain()
  renderTasks(tasks)
}

init()
