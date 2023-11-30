import React, { useState, useEffect } from 'react';
import './styles.css'
import logo from '$lib/media/Logo.svg';

const Header = ({ user }) => {
    useEffect(() => {
    }, [user]);
    return (
        <section id='Inicio' className="header-container">
            <header className="header">
                <div className="header-inner">
                    <Logo />
                    <NavBar />
                    <AuthButtons user={user} />
                </div>
            </header>
        </section>
    );
};

const Logo = () => {
    return (
        <div className="logo-container">
            <div className="logo-inner">
                <div className="logo-content">
                    <a href="/" className="logo-link">
                        <img src={logo} className="logo-image" alt="Avada Logo" />
                    </a>
                </div>
            </div>
        </div>
    );
}

const NavBarLi = ({ text, showArrow }) => {
    const [isHovered, setIsHovered] = useState(false);

    const handleMouseEnter = () => setIsHovered(true);
    const handleMouseLeave = () => setIsHovered(false);

    return (
        <li className={`nav-item ${isHovered ? 'x-safe-text' : ''}`}
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}>
            <a href={`#${text}`} className={`nav-link-text ${isHovered ? 'x-safe-text' : ''}`}>
                {text}
                {showArrow && <NavArrow />}
            </a>
        </li>
    );
};

const NavBar = () => {
    const menuItems = [
        { text: 'Inicio', showArrow: false },
        { text: 'Nuestros Objetivos', showArrow: false },
        { text: 'Acerca de nosotros', showArrow: false },
    ];

    return (
        <div className="nav-container">
            <div className="nav-inner-container">
                <nav className="nav-bar">
                    <ul className="nav-list">
                        {menuItems.map((item, index) => (
                            <NavBarLi key={index} text={item.text} showArrow={item.showArrow} />
                        ))}
                    </ul>
                </nav>
            </div>
        </div>
    );
};

const NavArrow = () => {
    return (
        <span className="nav-link-arrow">
            <svg fill="#ffffff" height="8px" width="8px" viewBox="0 0 330 330">
                <path fill="white" d="M325.607,79.393c-5.857-5.857-15.355-5.858-21.213,0.001l-139.39,139.393L25.607,79.393 
          c-5.857-5.857-15.355-5.858-21.213,0.001c-5.858,5.858-5.858,15.355,0,21.213l150.004,150
          c2.813,2.813,6.628,4.393,10.606,4.393s7.794-1.581,10.606-4.394l149.996-150
          C331.465,94.749,331.465,85.251,325.607,79.393z">
                </path>
            </svg>
        </span>
    );
};

const LogButton = ({ buttonText, link, extraClass, onMouseEnter, onMouseLeave, isTextShining }) => {

    const buttonClass = `cta-link ${extraClass}`;
    const textClass = `cta-text ${isTextShining ? 'x-safe-text' : ''}`;
    return (
        <a href={link}
            className={buttonClass}
            onMouseEnter={onMouseEnter}
            onMouseLeave={onMouseLeave}>
            <span className={textClass}>{buttonText}</span>
        </a>
    );
};
const AuthButtons = ({ user }) => {
    const [hoveredButton, setHoveredButton] = useState(null);
    const innerAlignStyle = {
        display: user ? "block" : "flex"
    };
    useEffect(() => {

    }, [user]);
    return (
        <div className="cta-container">
            <div className="cta-inner">
                <div className="cta-inner-align" style={innerAlignStyle}>

                    {user ? (
                        <LogButton
                            buttonText="Archivos"
                            link="/app"
                            extraClass={`cta-signup ${hoveredButton === 'login' ? 'cta-link-hover' : ''}`}
                            onMouseEnter={() => setHoveredButton('signup')}
                            onMouseLeave={() => setHoveredButton(null)}
                            isTextShining={hoveredButton === 'signup'}
                        />
                    ) : (
                        <>
                            <LogButton
                                buttonText="Registrar"
                                link="/register"
                                extraClass={hoveredButton === 'signup' ? 'cta-signup-hover' : ''}
                                onMouseEnter={() => setHoveredButton('login')}
                                onMouseLeave={() => setHoveredButton(null)}
                                isTextShining={hoveredButton === 'login'}
                            />
                            <LogButton
                                buttonText="Iniciar Sesion"
                                link="/login"
                                extraClass={`cta-signup ${hoveredButton === 'login' ? 'cta-link-hover' : ''}`}
                                onMouseEnter={() => setHoveredButton('signup')}
                                onMouseLeave={() => setHoveredButton(null)}
                                isTextShining={hoveredButton === 'signup'}
                            />
                        </>
                    )}
                </div>
            </div>
        </div>
    );
};

export default Header;
