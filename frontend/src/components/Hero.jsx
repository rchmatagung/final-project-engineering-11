import React from 'react'
import heroImg from '../assets/cbg.png'

const Hero = () => {
  return (
    <div name='home' className='w-full max-h-[1024px] bg-zinc-200 flex flex-col justify-between'>
        <div className='bg-white grid md:grid-cols-2 max-w-[1024px] mx-8'>
            <div className='flex flex-col justify-center md:items-start w-full px-2 '>
                <h1 className=' text-4xl md:text-5xl lg:text-7xl font-bold'>Let's find <span className="text-yellow-500 ">your Mentor</span></h1>
                <h1 className=' text-3xl md:text-5xl lg:text-7xl font-bold'>Build <span className="text-yellow-500">your code</span></h1>
                <p className='text-sm md:text-base lg:text-xl text-gray-600'>Chat mentor, ajukan belajar private , tanya jawab seputaran programming, update informasi seputar programming, semua bisa di Hicoder !</p>
            </div>
            <div>
                <img className='w-full' src={heroImg} alt="/" />
            </div>
        </div>
    </div>
  )
}

export default Hero