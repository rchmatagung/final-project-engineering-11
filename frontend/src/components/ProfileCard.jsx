import React from 'react'
import logoImg from '../assets/logohi.png'
import UpdateProfile from './UpdateProfile'
// import api from "./api/api"

const ProfileCard = () => {
    // const [user, setUser] = useState({})

    // const getUser = async () => {
    //     try {
    //       api.get('api/user/profile');
    //       then((res) => {
    //         console.log(res)
    //     })
    //     } catch (error) {
    //       console.log(error);
    //     }
    // };
    
    // useEffect(() => {
    //     getUser();
    // }, []);

    return (
        <div className='bg-zinc-200 flex flex-col justify-between'>
            <div className='bg-gray-200 grid md:grid-cols-2 mx-20'>
                <div className='bg-gray-500 m-5'>
                    <img src={logoImg} alt="user-profile-pic" />
                </div>
                <div className='flex flex-col justify-center md:items-start w-full px-2 m-5'>
                    <h1 className='py-3 text-4xl md:text-5xl lg:text-7xl font-bold'>Profile Data</h1>
                    <UpdateProfile />
                    <table>
                        <tr>
                            <td><p className='text-sm md:text-base lg:text-xl font-bold'>Username</p></td>
                            <td><p className='text-sm md:text-base lg:text-xl'>Username kamu</p></td>
                        </tr>
                        <tr>
                            <td><p className='text-sm md:text-base lg:text-xl font-bold'>Nama</p></td>
                            <td><p className='text-sm md:text-base lg:text-xl'>Nama kamu</p></td>
                        </tr>
                        <tr>
                            <td><p className='text-sm md:text-base lg:text-xl font-bold'>Address</p></td>
                            <td><p className='text-sm md:text-base lg:text-xl'>Address kamu</p></td>
                        </tr>
                        <tr>
                            <td><p className='text-sm md:text-base lg:text-xl font-bold'>No HP</p></td>
                            <td><p className='text-sm md:text-base lg:text-xl'>Nomor HP kamu</p></td>
                        </tr>
                        <tr>
                            <td><p className='text-sm md:text-base lg:text-xl font-bold'>Email</p></td>
                            <td><p className='text-sm md:text-base lg:text-xl'>Email kamu</p></td>
                        </tr>
                    </table>
                </div>
            </div>
        </div>
    )
}

export default ProfileCard