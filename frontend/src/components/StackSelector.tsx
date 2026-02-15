import { useEffect, useState } from "react"

interface Stack {
    id: number
    data: string
}

interface Props {
    value: string[]
    onChange: (v: string[]) => void
}

export default function StackSelector({ value, onChange }: Props) {
    const [options, setOptions] = useState<Stack[]>([])

    useEffect(() => {
        const cached = localStorage.getItem("stacks")
        if (cached) {
            // eslint-disable-next-line react-hooks/set-state-in-effect
            setOptions(JSON.parse(cached))
            return
        }

        fetch(`${import.meta.env.VITE_BACKEND_URL}/stacks`)
            .then((r) => r.json())
            .then((data) => {
                localStorage.setItem("stacks", JSON.stringify(data))
                setOptions(data)
            })
    }, [])

    return (
        <div>
            <p>Stack:</p>
            <select
                value={value}
                multiple
                onChange={(e) => {
                    const selected = Array.from(e.target.selectedOptions).map(
                        (opt) => opt.value
                    )
                    onChange(selected)
                }}
            >
                {options.map((o) => (
                    <option key={o.id} value={o.data}>
                        {o.data}
                    </option>
                ))}
            </select>
        </div>
    )
}
