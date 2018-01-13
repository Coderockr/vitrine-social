import React from 'react'
import {
  Container,
  Column,
  Columns,
  Nav,
  NavItem,
  NavGroup,
  Section,
  Title,
} from 're-bulma'
import './style.css';

const Header = () => {
  return (
    <Section className="header">
      <Nav className="nav">
        <Container>
          <NavGroup align="left">
            <NavItem className="navItem">
              <h3>
                <a href="#">
                  Vitrine Social
                </a>
              </h3>
            </NavItem>
          </NavGroup>
          <NavGroup align="right" isMenu>
            <NavItem className="navItem">
              <a href="#">
                Sobre o Projeto
              </a>
            </NavItem>
            <NavItem className="navItem">
              <a href="#">
              Quero Participar
              </a>
            </NavItem>
            <NavItem className="navItem">
              <a href="#">
                Contato
              </a>
            </NavItem>
          </NavGroup>
        </Container>
      </Nav>
    </Section>
  )
}

export default Header;
