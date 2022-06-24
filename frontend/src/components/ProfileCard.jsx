import React, {useState, useEffect} from 'react'
import logoImg from '../assets/logohi.png'
import UpdateProfile from './UpdateProfile'
import api from "../api/api"

const ProfileCard = () => {
    const [user, setUser] = useState({})

    const getUser = async () => {
        try {
        await api.get('/user/profile')
         .then((res) => {
            setUser(res.data.data)
            console.log(res.data.data)
        })
        } catch (error) {
          console.log(error);
        }
    };
    
    useEffect(() => {
        getUser();
    }, []);

    return (
        <div className='bg-zinc-200 flex flex-col justify-between'>
            <div className='bg-gray-200 grid md:grid-cols-2 mx-20'>
                <div className='bg-gray-500 m-5'>
                    <img src={logoImg} alt="user-profile-pic" />
                </div>
                <div className='flex flex-col justify-center md:items-start w-full px-2 m-5'>
                    <h1 className='py-3 text-4xl md:text-5xl lg:text-7xl font-bold'>Profile Data</h1>
                    <UpdateProfile 
                        refetch={getUser}
                    />
                    <table>
                        <tbody>
                            <tr>
                                <td><p className='text-sm md:text-base lg:text-xl font-bold'>Username</p></td>
                                <td><p className='text-sm md:text-base lg:text-xl pl-2'>{user.username}</p></td>
                            </tr>
                            <tr>
                                <td><p className='text-sm md:text-base lg:text-xl font-bold'>Nama</p></td>
                                <td><p className='text-sm md:text-base lg:text-xl pl-2'>{user.name}</p></td>
                            </tr>
                            <tr>
                                <td><p className='text-sm md:text-base lg:text-xl font-bold'>Address</p></td>
                                <td><p className='text-sm md:text-base lg:text-xl pl-2'>{user.address}</p></td>
                            </tr>
                            <tr>
                                <td><p className='text-sm md:text-base lg:text-xl font-bold'>No HP</p></td>
                                <td><p className='text-sm md:text-base lg:text-xl pl-2'>{user.phone}</p></td>
                            </tr>
                            <tr>
                                <td><p className='text-sm md:text-base lg:text-xl font-bold'>Email</p></td>
                                <td><p className='text-sm md:text-base lg:text-xl pl-2'>{user.email}</p></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    )
}

export default ProfileCard