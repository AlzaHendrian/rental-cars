export const ButtonSubmit = ({title, onClick}) => {
    return (
        <>
            <button 
            className=" bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:shadow-outline"
            onClick={onClick}
            >
                    {title}
            </button>
        </>
    )
}