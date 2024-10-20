export default function getResponseMessage(status: number, submessage: string): string {
  switch (status) {
    case 400:
      return 'Solicitação inválida. Verifique os dados informados.';
    case 401:
      return 'Não autorizado. Verifique suas credenciais.';
    case 403:
      return 'Acesso proibido.';
    case 404:
      return `${submessage} não encontrado(a).`;
    case 409:
      return `${submessage} já existe.`;
    case 500:
      return 'Erro interno no servidor. Tente novamente mais tarde.';
    case 502:
      return 'Erro ao se conectar ao servidor. Tente novamente.';
    case 503:
      return 'Serviço indisponível. O servidor está temporariamente fora do ar.';
    case 504:
      return 'Tempo de resposta excedido. O servidor está demorando para responder.';
    default:
      return `Erro desconhecido (Status: ${status}).`;
  }
}
