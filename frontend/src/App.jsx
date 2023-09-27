import { useEffect, useState } from "react"
import Navigation from "./navigation/Navigation"
import { API } from "./config/Api"
import { UserContextProvider } from "./context/UserContext"
import { BrowserRouter as Router } from "react-router-dom"

const App = () => {
  const [data, setData] = useState([])

  // get data form api
  const getData = async() => {
    try {
      const resp = await API.get("/cars")
      console.log(resp.data.data);
      setData(resp.data.data);
    }catch (err) {
      console.log(JSON.stringify(err, null, 4));
    }
  }

  useEffect(() => {
    getData()
  }, [])

  return (
    <Router>
      <UserContextProvider>
          <Navigation 
          dataCar={data}
          getData={getData}
          />
      </UserContextProvider>
    </Router>
  )
}

export default App