const maskPhone = (phone) => {
  const specialNumbers = /^(\(?(03|08|09)|(\(?40)\)?(\s)?(0))/;
  const capitals = /^\(?40[^\d]?0/;

  if (specialNumbers.test(phone)) {
    if (capitals.test(phone)) {
      return phone.replace(/\D/g, '')
        .replace(/(\d{8})(\d)/, '$1')
        .replace(/(\d{4})(\d)/, '$1-$2');
    }

    return phone.replace(/\D/g, '')
      .replace('(', '')
      .replace(')', '')
      .replace(' ', '')
      .replace(/(\d{11})(\d)/, '$1')
      .replace(/(\d{7})(\d)/, '$1-$2')
      .replace(/(\d{4})(\d{1,3})/, '$1 $2');
  }

  return phone.replace(/\D/g, '')
    .replace(/^(\d)/, '($1')
    .replace(/^(\(\d{2})(\d)/, '$1) $2')
    .replace(/(\d{4})(\d{1,4})/, '$1-$2')
    .replace(/(\d{5})(\d{5})/, '$1-$2')
    .replace(/(-\d{5})\d+?$/, '$1')
    .replace(/(\d{4})-(\d{1})(\d{4})/, '$1$2-$3');
};

const maskCpf = cpf => (
  cpf.replace(/\D/g, '')
    .replace(/(\d{3})(\d)/, '$1.$2')
    .replace(/(\d{3})(\d)/, '$1.$2')
    .replace(/(\d{3})(\d{1,2})$/, '$1-$2')
);

const maskCnpj = cnpj => (
  cnpj.replace(/\D/g, '')
    .replace(/^(\d{2})(\d)/, '$1.$2')
    .replace(/^(\d{2})\.(\d{3})(\d)/, '$1.$2.$3')
    .replace(/\.(\d{3})(\d)/, '.$1/$2')
    .replace(/(\d{4})(\d{1,2})$/, '$1-$2')
);

const maskCpfCnpj = (value) => {
  let identification = value;

  if (identification.length <= 14) {
    identification = maskCpf(identification);
  } else {
    identification = maskCnpj(identification);
  }
  return identification;
};

const maskCep = cep => (
  cep.replace(/(\d{5})(\d)/, '$1-$2')
);

const maskWhatsapp = whatsapp => (
  whatsapp.replace(/\D/g, '')
    .replace(/^(\d)/, '($1')
    .replace(/^(\(\d{2})(\d)/, '$1) $2')
    .replace(/(\d{4})(\d{1,4})/, '$1-$2')
    .replace(/(\d{5})(\d{5})/, '$1-$2')
    .replace(/(-\d{5})\d+?$/, '$1')
    .replace(/(\d{4})-(\d{1})(\d{4})/, '$1$2-$3')
);

export { maskPhone, maskCpf, maskCnpj, maskCpfCnpj, maskCep, maskWhatsapp };
