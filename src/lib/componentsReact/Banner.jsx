import './banner.css'
import React, { useEffect, useRef, createContext, useState } from 'react';
import x from '$lib/media/x.png'
import { gsap } from 'gsap'
import { useAnimationContext } from '$lib/componentsReact/hooks/useAnimation';
import NeonSection from './NeonSection';
import ObjectivesSection from './ObjectivesSection';
import AboutSection from './AboutSection';
import aboutImg from '$lib/media/about.png'
import Footer from './Footer';
import Carrusel from './Carrusel';

const AnimationContext = createContext();

export const AnimationProvider = ({ children }) => {
    const [isTitleAnimated, setTitleAnimated] = useState(false);

    return (
        <AnimationContext.Provider value={{ isTitleAnimated, setTitleAnimated }}>
            {children}
        </AnimationContext.Provider>
    );
};


const BannerTitle = React.forwardRef(({ children, shining }, ref) => {
    const className = `banner-title ${shining ? 'x-safe-text' : ''}`;

    return <h1 className={className} ref={ref}>{children}</h1>;
});

const BannerSubtitle = React.forwardRef(({ children }, ref) => (
    <p ref={ref} className="banner-subtitle">{children}</p>
));
const BannerButton = React.forwardRef(({ children, href }, ref) => {
    const [isHovered, setIsHovered] = useState(false);

    const handleMouseEnter = () => setIsHovered(true);
    const handleMouseLeave = () => setIsHovered(false);

    const shineClass = `${isHovered ? 'x-safe-text' : ''}`;
    const svgClass = `button-arrow`;

    return (
        <a
            ref={ref}
            className='banner-button'
            href={href}
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}
        >
            <span className={shineClass}>
                {children}
            </span>
            <svg className={svgClass} style={{ margin: "0 0 0 10px" }} width="24" height="24" xmlns="http://www.w3.org/2000/svg">
                <defs>
                    <linearGradient id="animatedGradient" x1="0%" y1="0%" x2="100%" y2="0%">
                        <stop offset="0%" stop-color="rgb(206, 46, 238)" id="startColor" />
                        <stop offset="50%" stop-color="rgb(15, 195, 233)" id="middleColor" />
                        <stop offset="100%" stop-color="rgb(206, 46, 238)" id="endColor" />
                    </linearGradient>
                </defs>
                <path d="M11.293 4.707 17.586 11H4v2h13.586l-6.293 6.293 1.414 1.414L21.414 12l-8.707-8.707-1.414 1.414z" fill="#000000" />
            </svg>

        </a>
    );
});


const BannerImg = () => {
    const imageRef = useRef(null);
    const divRef = useRef(null);
    const { isTitleAnimated } = useAnimationContext();
    useEffect(() => {

        const tl = gsap.timeline()
        if (isTitleAnimated) {
            tl.fromTo(
                divRef.current,
                1,
                { width: "0px", },
                { width: "45.7375em" }
            )
                .fromTo(
                    divRef.current,
                    1,
                    { height: "0px", },
                    { height: "24.9375em" }
                )
        }


    }, [isTitleAnimated]);
    return (
        <div className="main-banner-image" ref={divRef}>
            <div className='main-banner-image-animation'>
                <img
                    ref={imageRef}
                    className="main-banner-image-inner"
                    src={x}
                    alt="Descripción"
                />
            </div>
        </div>
    );
};

const ScrollToTop = () => {
    const scrollToTop = () => {
        const section = document.getElementById('Inicio');
        section.scrollIntoView({ behavior: 'smooth', block: 'start' });
    };

    return (
        <button onClick={scrollToTop} style={scrollToTopStyle}>
            ⬆️
        </button>
    );
};

const scrollToTopStyle = {
    position: 'fixed',
    bottom: '20px',
    right: '20px',
    background: 'none',
    border: 'none',
    fontSize: '2em',
    cursor: 'pointer',
    zIndex: 1000
};

const MainBanner = () => {
    const headlineRefs = useRef([React.createRef(), React.createRef(), React.createRef()]);
    headlineRefs.current = [];

    const { setTitleAnimated } = useAnimationContext();

    useEffect(() => {
        const tl = gsap.timeline()
        console.log(headlineRefs.current)
        headlineRefs.current.forEach((headline, index) => {
            tl.fromTo(
                headline,
                { opacity: 0, x: -20 },
                { opacity: 1, x: 0, duration: 0.3, delay: index * 0.2 }
            );
        });
        tl.add(() => {
            setTitleAnimated(true);
        });
    }, []);

    return (
        <>
            <section className='banner'>
                <div className="main-banner">
                    <div className='main-banner-inner'>
                        <div className='main-banner-content'>
                            <div className='main-banner-text'>
                                <BannerTitle ref={el => headlineRefs.current[0] = el}>Bienvenido a</BannerTitle>
                                <BannerTitle ref={el => headlineRefs.current[1] = el} shining={true}>X-safe</BannerTitle>
                                <BannerSubtitle ref={el => headlineRefs.current[2] = el}>
                                    ¡Aquí, nos dedicamos a ofrecerte información para asegurar tu experiencia en línea.
                                    Nuestro repositorio es el aliado perfecto para mantenerte al tanto de cualquier
                                    vulnerabilidad que pueda afectar tu presencia en X.
                                </BannerSubtitle>
                                <BannerButton ref={el => headlineRefs.current[3] = el} href="/app">
                                    Comenzar a contribuir
                                </BannerButton>
                            </div>
                        </div>
                        <BannerImg />
                    </div>
                </div>
            </section>
            <NeonSection />
            <ObjectivesSection />
            <Carrusel />
            <AboutSection
                title="Acerca de Nosotros"
                subtitle={"X-Safe: Comprometidos con la Investigación y la Seguridad en Línea"}
                email={"Xsafe@gmail.com"}
                imageSrc={aboutImg}
                emailText={"Contacto"}
            />
            <ScrollToTop />
            <Footer
                aboutTitle="Acerca de Nosotros"
                aboutText="X-Safe es un proyecto dedicado a fortalecer la seguridad en Twitter."
                contactInfo="Información de contacto:"
                phoneNumber="+58 (424) 8024132"
                email="xsafe@gmail.com"
                authorInfo="Este es un prototipo con respecto al proyecto de grado. Br Bárbara Barreca"
            />
        </>
    );
}


export default MainBanner;
