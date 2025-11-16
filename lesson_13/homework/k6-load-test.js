import http from 'k6/http';

export default function () {
	const url = 'http://localhost:3000/user/search?first_name="Але"&second_name="Пет"';

	http.get(url);
}
