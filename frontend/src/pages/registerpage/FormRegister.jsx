import { Link } from "react-router-dom"
import { ButtonSubmit } from "../../components/ButtonSubmit"
import { Input } from "../../components/Input"

const FormRegister = ({register, form, handleOnChange}) => {
    const data = [
        {id: 1, placeHolderRegister:'input your name', param: 'name', value: form.name},
        {id: 2, placeHolderRegister:'input your email', param: 'email', value: form.email},
        {id: 3, placeHolderRegister:'input your password', param: 'password', value: form.password},
    ]
    return (
        <div className="w-[70vh] shadow-xl p-5 rounded-md mx-auto mt-20">
            <h1 className="text-2xl font-semibold mb-3">Sign Up</h1>
            {data.map((item) => (
                <div key={item.id} className="mb-2">
                    <Input 
                    placeholder={item.placeHolderRegister}
                    value={item.value}
                    onChange={(e) => handleOnChange(item.param, e.target.value)}
                    />
                </div>
            ))}
            <div className="grid grid-cols-1 mt-5">
                <ButtonSubmit
                title="sign Up"
                onClick={register}/>
            </div>
            <p className="text-sm text-center mt-5">
                Already have an account? click{' '}
                <Link className="font-semibold cursor-pointer" to={'/login'}>
                    Here
                </Link>
            </p>
        </div>
    )
}

export default FormRegister