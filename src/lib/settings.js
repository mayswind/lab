import Cookies from 'js-cookie';

import statisticsConstants from '../consts/statistics.js';

const settingsLocalStorageKey = 'lab_app_settings';
const serverSettingsCookieKey = 'lab_server_settings';

const defaultSettings = {
    lang: 'en',
    debug: false,
    applicationLock: false,
    applicationLockWebAuthn: false,
    autoUpdateExchangeRatesData: true,
    thousandsSeparator: true,
    currencyDisplayMode: 'symbol', // or 'none' or 'code' or 'name'
    showAmountInHomePage: true,
    showAccountBalance: true,
    statistics: {
        defaultChartType: statisticsConstants.defaultChartType,
        defaultChartDataType: statisticsConstants.defaultChartDataType,
        defaultDataRangeType: statisticsConstants.defaultDataRangeType,
        defaultAccountFilter: {},
        defaultTransactionCategoryFilter: {},
        sortingType: statisticsConstants.defaultSortingType
    },
    animate: true,
    autoDarkMode: true
};

function getOriginalSettings() {
    try {
        const storageData = localStorage.getItem(settingsLocalStorageKey) || '{}';
        return JSON.parse(storageData);
    } catch (ex) {
        console.warn('settings in local storage is invalid', ex);
        return {};
    }
}

function getFinalSettings() {
    return Object.assign({}, defaultSettings, getOriginalSettings());
}

function setSettings(settings) {
    const storageData = JSON.stringify(settings);
    return localStorage.setItem(settingsLocalStorageKey, storageData);
}

function getOriginalOption(key) {
    return getOriginalSettings()[key];
}

function getOption(key) {
    return getFinalSettings()[key];
}

function getSubOption(key, subKey) {
    const options = getFinalSettings()[key] || {};
    return options[subKey];
}

function setOption(key, value) {
    if (!Object.prototype.hasOwnProperty.call(defaultSettings, key)) {
        return;
    }

    const settings = getFinalSettings();
    settings[key] = value;

    return setSettings(settings);
}

function setSubOption(key, subKey, value) {
    if (!Object.prototype.hasOwnProperty.call(defaultSettings, key)) {
        return;
    }

    if (!Object.prototype.hasOwnProperty.call(defaultSettings[key], subKey)) {
        return;
    }

    const settings = getFinalSettings();
    let options = settings[key];

    if (!options) {
        options = {};
    }

    options[subKey] = value;
    settings[key] = options;

    return setSettings(settings);
}

function getServerSetting(key) {
    const settings = Cookies.get(serverSettingsCookieKey) || '';
    const settingsArr = settings.split('_');

    for (let i = 0; i < settingsArr.length; i++) {
        const pairs = settingsArr[i].split('.');

        if (pairs[0] === key) {
            return pairs[1];
        }
    }

    return undefined;
}

function clearSettings() {
    localStorage.removeItem(settingsLocalStorageKey);
}

export default {
    getLanguage: () => getOriginalOption('lang'),
    setLanguage: value => setOption('lang', value),
    isEnableDebug: () => getOption('debug'),
    setEnableDebug: value => setOption('debug', value),
    isEnableApplicationLock: () => getOption('applicationLock'),
    setEnableApplicationLock: value => setOption('applicationLock', value),
    isEnableApplicationLockWebAuthn: () => getOption('applicationLockWebAuthn'),
    setEnableApplicationLockWebAuthn: value => setOption('applicationLockWebAuthn', value),
    isAutoUpdateExchangeRatesData: () => getOption('autoUpdateExchangeRatesData'),
    setAutoUpdateExchangeRatesData: value => setOption('autoUpdateExchangeRatesData', value),
    isEnableThousandsSeparator: () => getOption('thousandsSeparator'),
    setEnableThousandsSeparator: value => setOption('thousandsSeparator', value),
    getCurrencyDisplayMode: () => getOption('currencyDisplayMode'),
    setCurrencyDisplayMode: value => setOption('currencyDisplayMode', value),
    isShowAmountInHomePage: () => getOption('showAmountInHomePage'),
    setShowAmountInHomePage: value => setOption('showAmountInHomePage', value),
    isShowAccountBalance: () => getOption('showAccountBalance'),
    setShowAccountBalance: value => setOption('showAccountBalance', value),
    getStatisticsDefaultChartType: () => getSubOption('statistics', 'defaultChartType'),
    setStatisticsDefaultChartType: value => setSubOption('statistics', 'defaultChartType', value),
    getStatisticsDefaultChartDataType: () => getSubOption('statistics', 'defaultChartDataType'),
    setStatisticsDefaultChartDataType: value => setSubOption('statistics', 'defaultChartDataType', value),
    getStatisticsDefaultDateRange: () => getSubOption('statistics', 'defaultDataRangeType'),
    setStatisticsDefaultDateRange: value => setSubOption('statistics', 'defaultDataRangeType', value),
    getStatisticsDefaultAccountFilter: () => getSubOption('statistics', 'defaultAccountFilter'),
    setStatisticsDefaultAccountFilter: value => setSubOption('statistics', 'defaultAccountFilter', value),
    getStatisticsDefaultTransactionCategoryFilter: () => getSubOption('statistics', 'defaultTransactionCategoryFilter'),
    setStatisticsDefaultTransactionCategoryFilter: value => setSubOption('statistics', 'defaultTransactionCategoryFilter', value),
    getStatisticsSortingType: () => getSubOption('statistics', 'sortingType'),
    setStatisticsSortingType: value => setSubOption('statistics', 'sortingType', value),
    isEnableAnimate: () => getOption('animate'),
    setEnableAnimate: value => setOption('animate', value),
    isEnableAutoDarkMode: () => getOption('autoDarkMode'),
    setEnableAutoDarkMode: value => setOption('autoDarkMode', value),
    isUserRegistrationEnabled: () => getServerSetting('r') === '1',
    isDataExportingEnabled: () => getServerSetting('e') === '1',
    clearSettings: clearSettings
};
