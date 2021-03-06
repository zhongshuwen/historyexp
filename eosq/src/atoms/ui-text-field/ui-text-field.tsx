import { styled } from "../../theme"
import Input from "@material-ui/core/Input"
import * as React from "react"
import Paper from "@material-ui/core/Paper/Paper"

export const UiInput: React.ComponentType<any> = styled(Input)`
  font-size: 16px !important;
  width: ${(props: any) => (props.width ? props.width : "250px")} !important;
  background-color: ${(props) => props.theme.colors.formFieldBg};
  padding-left: 8px;
  box-sizing: border-box;
  height: 35px;

  input:-webkit-autofill {
    background-color: ${(props) => props.theme.colors.formFieldBg} !important;
  }

  input:-webkit-autofill {
    -webkit-text-fill-color: black;
    -webkit-box-shadow: 0 0 0 30px ${(props) => props.theme.colors.formFieldBg} inset;
  }
`

export const UiSearch: React.ComponentType<any> = styled(Input)`
  background-color: ${(props: any) => "#fff"};
  border:1px solid transparent;
  border-bottom:3px solid #aaa;
  color: #222 !important;
  padding: 0.75em 95px 0.75em 25px;
  font-size: 18px !important;
  height: 76px;
  width: 100%;
  outline: none;
  text-align: left;
  font-family: ${(props) => props.theme.fontFamily.robotoCondensed} !important;
  border-radius: 0px !important;
  transition: background-color 500ms linear, border-color 500ms linear;

  input:-webkit-autofill {
    color: #222 !important;
    background-color: ${(props) => props.theme.colors.searchBg} !important;
  }

  input:-webkit-autofill {
    -webkit-text-fill-color: #222;
    -webkit-box-shadow: 0 0 0 30px ${(props) => props.theme.colors.searchBg} inset;
  }

  &:hover{
    background-color: #efefef !important;
    border-bottom-color: #000;
  }
  &:focus {
    border-color: #000;
    border-bottom-color: #000;
    background-color: #efefef !important;
  }

  input::placeholder {
    opacity: 1;
    font-size: 18px !important;
    font-family: ${(props) => props.theme.fontFamily.robotoCondensed} !important;

    font-weight: 400;
    color: ${(props) => props.theme.colors.bleu5} !important;
  }

  @media (max-width: 767px) {
    height: 56px;
    padding: 0.75em 43px 0.75em 55px;
  }
`

export const UiPaper: React.ComponentType<any> = styled(Paper)`
  position: absolute;
  z-index: 1;
  margin-top: 0;
  left: 0;
  right: 0;
`
