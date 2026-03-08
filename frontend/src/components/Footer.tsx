import styles from "./Footer.module.css"

export default function Footer() {
    return (
        <footer className={styles.footer}>
            <small>Go Predict · Built by Rafael · <a href="https://github.com/rafaeldepontes/go-predict" target="_blank"
                rel="noopener" style={{ color: "var(--accent-color)" }}>Open on GitHub</a></small>
        </footer>
    )
}