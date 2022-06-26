import React from 'react'
import heroImg from '../assets/cbg.png'

const Hero = () => {
  return (
    <div name='home' className='w-full bg-white flex flex-col justify-between'>
        <div className='bg-white grid md:grid-cols-2 max-w-[1024px] mx-auto'>
            <div className='flex flex-col justify-center md:items-start w-full pt-5 px-10'>
                <h1 className=' text-5xl lg:text-6xl font-bold'>Let's find <span className="text-yellow-500 ">your Mentor</span></h1>
                <h1 className=' text-5xl lg:text-6xl font-bold'>Build <span className="text-yellow-500">your code</span></h1>
                <p className='text-sm md:text-base lg:text-xl text-gray-600 pt-3'>Chat mentor, ajukan belajar private , tanya jawab seputaran programming, update informasi seputar programming, semua bisa di Hicoder !</p>
            </div>
            <div>
                <img className='w-full px-20 md:w-full lg:px-0' src={heroImg} alt="/" />
            </div>
        </div>
    </div>
  )
}

export default Hero