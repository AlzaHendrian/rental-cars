import { useState } from "react";
import FormRegister from "./FormRegister"
import { API } from "../../config/Api";
import { useNavigate } from "react-router-dom";

const RegisterPage = () => {
    const navigate = useNavigate()
    const [form, setForm] = useState({
        name: '',
        email: '',
        password: '',
    });
    const handleOnChange = (name, value) => {
        setForm({
            ...form,
            [name]: value
        })
    };

    const register = async () => {
        try {
            const response = await API.post('/register', form);
            console.log(response);
            setForm({
                name: '',
                email: '',
                password: '',
            })
            alert("register success")
            
            // status check
            if (response.data.data) {
              navigate('/login');
            } else {
              navigate('/register');
            }
            console.log(response.data.data)
      
          } catch (err) {
            console.log(JSON.stringify(err, null, 2));
          }
    }
    return (
        <>
        <FormRegister
        form={form}
        register={register}
        handleOnChange={handleOnChange}
        />
        </>
    )
}

export default RegisterPage