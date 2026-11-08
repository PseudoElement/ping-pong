package sounds

type AppSound string

/**
 * AppSound should start from "../" then it normalized for build dirname
 */
const (
	CoinEarn AppSound = "../assets/audios/coin-earn.mp3"
)

var ALL_SOUNDS []AppSound = []AppSound{
	CoinEarn,
}
