interface Props {
    value: number
    onChange: (v: number) => void
}

export default function TeamSizeTextBox({ value, onChange }: Props) {
    return (
        <div>
            <p>Team Size:</p>
            <input
                type="number"
                min={1}
                value={value}
                onChange={(e) => {
                    let value = Number(e.target.value)

                    if (value != null && value < 1) {
                        value *= -1
                    }

                    onChange(value)
                }}
            />
        </div>
    )
}
