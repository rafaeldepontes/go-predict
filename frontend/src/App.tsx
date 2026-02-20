import { useState } from "react"
import "./App.css"

import FeatureTextBox from "./components/FeatureTextBox"
import PredictionField from "./components/PredictionField"
import StackSelector from "./components/StackSelector"
import TeamSizeTextBox from "./components/TeamSizeTextBox"
import TeamSkillSelector from "./components/TeamSkillSelector"
import ErrorPopUp from "./components/ErrorPopUp"

interface Response {
	data: string,
}

interface RequestError {
	occurred: boolean,
	msg: string,
}

const initialError: RequestError = { msg: "", occurred: false }

export default function App() {
	const [teamSize, setTeamSize] = useState(1)
	const [seniority, setSeniority] = useState("")
	const [stack, setStack] = useState<string[]>([])
	const [feature, setFeature] = useState("")
	const [prediction, setPrediction] = useState("")
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState<RequestError>(initialError)

	const handlePredict = async () => {
		setLoading(true)
		setPrediction("")
		setError(initialError)

		try {
			let stackReq = stack.join(" ")

			const resp = await fetch(`${import.meta.env.VITE_BACKEND_URL}/predict`, {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({
					body: feature,
					size: teamSize,
					stack: stackReq,
					level: seniority,
				}),
			})

			if (!resp.ok) {
				throw new Error(await resp.text())
			}

			const text: Response = await resp.json()
			setPrediction(text.data)
		} catch (err) {
			if (err instanceof Error) {
				setError({
					msg: err.message,
					occurred: true,
				})
			} else {
				setError({
					msg: "Something went wrong",
					occurred: true,
				})
			}
		} finally {
			setLoading(false)
		}
	}

	return (
		<>
			{error.occurred && <ErrorPopUp error={error.msg} onClose={() => setError(initialError)} />}
			<main className="container">
				<header style={{ gridColumn: '1 / -1' }}>
					<h1>Project Estimator<span style={{ color: 'var(--accent-color)' }}>.</span></h1>
					<p style={{ color: 'var(--text-secondary)', marginTop: '-10px' }}>
						Put your team and predict effort.
					</p>
				</header>

				<div className="left">
					<div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '12px' }}>
						<TeamSizeTextBox value={teamSize} onChange={setTeamSize} />
						<TeamSkillSelector value={seniority} onChange={setSeniority} />
					</div>

					<StackSelector value={stack} onChange={setStack} />
					<FeatureTextBox value={feature} onChange={setFeature} />

					<button onClick={handlePredict} disabled={loading}>
						{loading ? "Calculating..." : "Make Prediction"}
					</button>
				</div>

				<div className="right">
					<PredictionField value={prediction} loading={loading} />
				</div>

				<footer className="footer">
					<small>Go Predict · Built by Rafael · <a href="https://github.com/rafaeldepontes/go-predict" target="_blank"
						rel="noopener" style={{ color: "var(--accent)" }}>Open on GitHub</a></small>
				</footer>
			</main>
		</>
	)
}