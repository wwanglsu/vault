import { helper as buildHelper } from '@ember/component/helper';

export const KMIP = {
  displayName: 'KMIP',
  value: 'kmip',
  type: 'kmip',
  category: 'generic',
};

const MOUNTABLE_SECRET_ENGINES = [
  {
    displayName: 'AliCloud KMS',
    value: 'alicloud',
    type: 'alicloud',
    category: 'cloud',
  },
  {
    displayName: 'AWS KMS',
    value: 'aws',
    type: 'aws',
    category: 'cloud',
  },
  {
    displayName: 'Azure',
    value: 'azure',
    type: 'azure',
    category: 'cloud',
  },
  {
    displayName: 'Google Cloud KMS',
    value: 'gcpkms',
    type: 'gcpkms',
    category: 'cloud',
  },
  {
    displayName: 'Secrets Storage',
    value: 'kv',
    type: 'kv',
    category: 'generic',
  },
  {
    displayName: 'VMC KMS',
    value: 'vmckms',
    type: 'transit',
    category: 'generic',
  },
];

export function engines() {
  return MOUNTABLE_SECRET_ENGINES.slice();
}

export default buildHelper(engines);
