import React from 'react'
import heroImg from '../../assets/cbg.png'

const Hero = () => {
  return (
    <div className='bg-zinc-200'>
        <div name='home' className='container px-6 py-4 mx-auto lg:flex lg:h-128 lg:py-16'>
            <div className='flex flex-col items-center w-full lg:flex-row lg:w-1/2'>
                <div className='flex flex-col justify-center md:items-start w-full px-2 '>
                    <h1 className='py-3 text-4xl md:text-5xl lg:text-5xl font-bold'>Let's find <span className="text-yellow-400">your Mentor</span></h1>
                    <h1 className='py-3 text-4xl md:text-5xl lg:text-5xl font-bold'>Build <span className="text-yellow-400">your Code</span></h1>
                    <p className='mt-5 text-gray-600'>Chat mentor, ajukan belajar private , tanya jawab seputaran programming, update informasi seputar programming, semua bisa di Hicoder !</p>
                </div>
            </div>
            <div className="flex items-center justify-center w-full mt-2 lg:h-96 lg:w-1/2">
                <img className='object-cover w-full max-w-2xl rounded-md lg:h-full' src={heroImg} alt="/" />
            </div>
        </div>
    </div>
  )
}

export default Hero